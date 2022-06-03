package service

import (
	"bankingApp/domain"
	"bankingApp/dto"
	"bankingApp/errs"
)

type AccountService interface {
	NewAccount(request dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
	NewTransaction(request dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (s DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	account := domain.NewAccount(req.CustomerId, req.AccountType, req.Amount)
	if newAccount, err := s.repo.Save(account); err != nil {
		return nil, err
	} else {
		return newAccount.ToNewAccountResponseDto(), nil
	}
}

func (s DefaultAccountService) NewTransaction(req dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errs.AppError) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	transaction := domain.NewTransaction(req.AccountId, req.TransactionType, req.Amount)
	if newTransaction, err := s.repo.Update(transaction); err != nil {
		return nil, err
	} else {
		return newTransaction.ToNewTransactionResponseToDto(), nil
	}
}

func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo}
}
