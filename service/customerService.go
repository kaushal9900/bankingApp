package service

import (
	"bankingApp/domain"
	"bankingApp/dto"
	"bankingApp/errs"
)

type CustomerService interface {
	GetAllCustomers(string) ([]*dto.CustomerReponse, *errs.AppError)
	GetCustomer(string) (*dto.CustomerReponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers(status string) ([]*dto.CustomerReponse, *errs.AppError) {
	c, err := s.repo.FindAll(status)
	if err != nil {
		return nil, err
	}
	var response []*dto.CustomerReponse
	for _, v := range c {
		response = append(response, v.ToDto())
	}
	return response, nil
}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerReponse, *errs.AppError) {
	c, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}
	response := c.ToDto()
	return response, nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
