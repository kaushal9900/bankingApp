package domain

import (
	"bankingApp/dto"
	"bankingApp/errs"
)

const dbTSLayout = "2006-01-02 15:04:05"

type Account struct {
	AccountId   string  `db:"account_id"`
	CustomerId  string  `db:"customer_id"`
	OpeningDate string  `db:"opening_date"`
	AccountType string  `db:"account_type"`
	Amount      float64 `db:"amount"`
	Status      string  `db:"status"`
}

type Transaction struct {
	TransactionId   string  `db:"amount float64"`
	AccountId       string  `db:"account_id"`
	Amount          float64 `db:"amount"`
	TransactionType string  `db:"transaction_type"`
	TransactionDate string  `db:"transaction_date"`
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
	Update(Transaction) (*Transaction, *errs.AppError)
}

func (a Account) ToNewAccountResponseDto() *dto.NewAccountResponse {
	return &dto.NewAccountResponse{a.AccountId}
}

func (t Transaction) ToNewTransactionResponseToDto() *dto.NewTransactionResponse {
	return &dto.NewTransactionResponse{t.TransactionId}
}

func NewAccount(customerId, accountType string, amount float64) Account {
	return Account{
		CustomerId:  customerId,
		OpeningDate: dbTSLayout,
		AccountType: accountType,
		Amount:      amount,
		Status:      "1",
	}
}

func NewTransaction(accountId, transactionType string, amount float64) Transaction {
	return Transaction{
		TransactionType: transactionType,
		Amount:          amount,
		AccountId:       accountId,
	}
}
