package domain

import (
	"bankingApp/errs"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDb struct {
	db *sqlx.DB
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	customers := make([]Customer, 0)
	var err error
	if status != "" {
		findAllSql := "select customer_id,name,city,zipcode,date_of_birth,status from customers where status = ?"
		err = d.db.Select(&customers, findAllSql, status)
	} else {
		findAllSql := "select customer_id,name,city,zipcode,date_of_birth,status from customers"
		err = d.db.Select(&customers, findAllSql)
	}
	if err != nil {
		log.Println(err)
		return nil, errs.NewUnexpectedError("unexpected db error")
	}
	// err = sqlx.StructScan(rows, &customers)
	// if err != nil {
	// 	logger.Error("Error while Scanning customer" + err.Error())
	// 	return nil, errs.NewUnexpectedError("unexpected db error")
	// }
	// for rows.Next() {
	// 	var c Customer
	// 	err := rows.Scan(&c.Id, &c.Name, &c.City, &c.ZipCode, &c.DateofBirth, &c.Status)
	// 	if err != nil {
	// 		if err == sql.ErrNoRows {
	// 			return nil, err
	// 		} else {
	// 			return nil, err
	// 		}
	// 	}
	// 	customers = append(customers, c)
	// }
	return customers, nil

}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	customerSql := "select customer_id,name,city,zipcode,date_of_birth,status from customers where customer_id = ?"

	//row := d.db.QueryRow(customerSql, id)
	//sqlx
	var c Customer
	err := d.db.Get(&c, customerSql, id)
	//err := row.Scan(&c.Id, &c.Name, &c.City, &c.ZipCode, &c.DateofBirth, &c.Status)
	if err != nil {
		log.Println(err)
		return nil, errs.NewNotFoundError("not Found")
	}
	return &c, nil
}

// func (d CustomerRepositoryDb) Insert() (*errs.AppError) {
// 	customerSql := "select customer_id,name,city,zipcode,date_of_birth,status from customers where customer_id = ?"

// }

func NewCustomerRepositoryDb(dbClient *sqlx.DB) CustomerRepositoryDb {
	return CustomerRepositoryDb{dbClient}
}
