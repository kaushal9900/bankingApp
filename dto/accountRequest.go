package dto

import (
	"bankingApp/errs"
	"strings"
)

type NewAccountRequest struct {
	CustomerId  string  `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

type NewTransactionRequest struct {
	AccountId       string  `json:"account_id"`
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transaction_type"`
}

func (r NewAccountRequest) Validate() *errs.AppError {
	if r.Amount < 5000 {
		return errs.NewValidationError("To Open new account u need to deposit at least 5000")
	}
	if strings.ToLower(r.AccountType) != "saving" && strings.ToLower(r.AccountType) != "checking" {
		return errs.NewValidationError("Account type needs to be saving or checking")
	}
	return nil
}

func (r NewTransactionRequest) Validate() *errs.AppError {
	if r.Amount < 1 {
		return errs.NewValidationError("Amount can't be negative")
	}
	if strings.ToLower(r.TransactionType) != "deposit" && strings.ToLower(r.TransactionType) != "withdraw" {
		return errs.NewValidationError("Transaction type should be deposit or withdraw")
	}
	return nil
}
