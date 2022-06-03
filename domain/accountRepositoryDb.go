package domain

import (
	"bankingApp/errs"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
)

type AccountRepositoryDb struct {
	db *sqlx.DB
}

func (d AccountRepositoryDb) Save(acc Account) (*Account, *errs.AppError) {
	accountInsertSql := "Insert INTO accounts" +
		"(customer_id, opening_date, account_type, amount, status)" +
		"values(?,?,?,?,?)"
	res, err := d.db.Exec(accountInsertSql, acc.CustomerId, acc.OpeningDate,
		acc.AccountType, acc.Amount, acc.Status)
	if err != nil {
		log.Fatal("Error while creating new account", err)
		return nil, errs.NewUnexpectedError("Unexpected error from mysql")
	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal("Error while getting last insert id for new account", res)
		return nil, errs.NewUnexpectedError("Unexpected error from mysql")
	}
	acc.AccountId = strconv.FormatInt(id, 10)
	return &acc, nil
}

func (d AccountRepositoryDb) Update(t Transaction) (*Transaction, *errs.AppError) {
	balanceSql := "SELECT amount FROM accounts WHERE account_id=?"
	//account := make([]Account, 0)
	var balance float64
	err := d.db.Get(&balance, balanceSql, t.AccountId)
	if err != nil {
		fmt.Println(err)
		return nil, errs.NewUnexpectedError("Unexpected error while getting balance")
	}
	updateBalancesql := "UPDATE accounts SET AMOUNT=? WHERE account_id=?"
	if strings.ToLower(t.TransactionType) == "deposit" {
		balance = balance + t.Amount
	} else {
		if balance-t.Amount < 0 {
			return nil, errs.NewValidationError("Insufficent balance to withdraw")
		}
		balance = balance - t.Amount
	}
	_, err = d.db.Exec(updateBalancesql, balance, t.AccountId)
	if err != nil {
		return nil, errs.NewUnexpectedError("Error while updating balance")
	}
	transSql := "INSERT INTO transactions" +
		"(account_id,amount,transaction_type)" +
		"values(?,?,?)"
	res, err := d.db.Exec(transSql, t.AccountId, t.Amount, t.TransactionType)
	if err != nil {
		return nil, errs.NewUnexpectedError("Error while inserting transaction")
	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal("Error while getting last insert id for new account", res)
		return nil, errs.NewUnexpectedError("Unexpected error from mysql")
	}
	t.TransactionId = strconv.FormatInt(id, 10)
	return &t, nil
}

func NewAccountRepositoryDb(dbclient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{dbclient}
}
