package domain

import (
	"bankingApp/dto"
	"bankingApp/errs"
)

type Customer struct {
	Id          string `db:"customer_id"`
	Name        string
	City        string
	ZipCode     string
	DateofBirth string `db:"date_of_birth"`
	Status      string
}

type CustomerRepository interface {
	FindAll(string) ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
}

func (c Customer) status() string {
	status := "active"
	if c.Status == "0" {
		status = "inactive"
	}
	return status
}

func (c Customer) ToDto() *dto.CustomerReponse {

	return &dto.CustomerReponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		ZipCode:     c.ZipCode,
		DateofBirth: c.DateofBirth,
		Status:      c.status(),
	}
}
