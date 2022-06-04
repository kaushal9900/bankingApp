package domain

import (
	"bankingApp/domain/logger"
	"bankingApp/errs"
	"log"
	"strconv"

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

// func (d AccountRepositoryDb) Update(t Transaction) (*Transaction, *errs.AppError) {
// 	balanceSql := "SELECT amount FROM accounts WHERE account_id=?"
// 	//account := make([]Account, 0)
// 	var balance float64
// 	err := d.db.Get(&balance, balanceSql, t.AccountId)
// 	if err != nil {
// 		fmt.Println(err)
// 		return nil, errs.NewUnexpectedError("Unexpected error while getting balance")
// 	}
// 	updateBalancesql := "UPDATE accounts SET AMOUNT=? WHERE account_id=?"
// 	if strings.ToLower(t.TransactionType) == "deposit" {
// 		balance = balance + t.Amount
// 	} else {
// 		if balance-t.Amount < 0 {
// 			return nil, errs.NewValidationError("Insufficent balance to withdraw")
// 		}
// 		balance = balance - t.Amount
// 	}
// 	_, err = d.db.Exec(updateBalancesql, balance, t.AccountId)
// 	if err != nil {
// 		return nil, errs.NewUnexpectedError("Error while updating balance")
// 	}
// 	transSql := "INSERT INTO transactions" +
// 		"(account_id,amount,transaction_type)" +
// 		"values(?,?,?)"
// 	res, err := d.db.Exec(transSql, t.AccountId, t.Amount, t.TransactionType)
// 	if err != nil {
// 		return nil, errs.NewUnexpectedError("Error while inserting transaction")
// 	}
// 	id, err := res.LastInsertId()
// 	if err != nil {
// 		log.Fatal("Error while getting last insert id for new account", res)
// 		return nil, errs.NewUnexpectedError("Unexpected error from mysql")
// 	}
// 	t.TransactionId = strconv.FormatInt(id, 10)
// 	return &t, nil
// }

func (d AccountRepositoryDb) SaveTransaction(t Transaction) (*Transaction, *errs.AppError) {
	tx, err := d.db.Begin()
	if err != nil {
		logger.Error("Error while starting new transaction")
		return nil, errs.NewUnexpectedError("Unexpected error while starting new transaction")
	}
	transactionQuery := "INSERT INTO transactions (account_id, amount, transaction_type, transaction_date) values (?, ?, ?, ?)"
	res, _ := tx.Exec(transactionQuery, t.AccountId, t.Amount, t.TransactionType, t.TransactionDate)

	if t.IsWithdraw() {
		_, err = tx.Exec(`UPDATE accounts SET amount = amount - ? where account_id = ?`, t.Amount, t.AccountId)
	} else {
		_, err = tx.Exec(`UPDATE accounts SET amount = amount + ? where account_id = ?`, t.Amount, t.AccountId)
	}
	// rollback in case of some errs
	if err != nil {
		tx.Rollback()
		logger.Error("Error while saving transaction" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected DB Error")
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.Error("Error while commiting transaction" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected DB Error")
	}
	transactionId, err := res.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last transaction id" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected DB Error")
	}
	account, appErr := d.FindBy(t.AccountId)
	if appErr != nil {
		return nil, appErr
	}
	t.TransactionId = strconv.FormatInt(transactionId, 10)
	t.Amount = account.Amount
	return &t, nil
}

func (d AccountRepositoryDb) FindBy(accountId string) (*Account, *errs.AppError) {
	findSql := "Select account_id,customer_id,opening_date,account_type,amount from accounts where account_id = ?"
	var account Account
	err := d.db.Get(&account, findSql, accountId)
	if err != nil {
		logger.Error("Error while fetching account details " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected DB Error")
	}
	return &account, nil
}

func NewAccountRepositoryDb(dbclient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{dbclient}
}
