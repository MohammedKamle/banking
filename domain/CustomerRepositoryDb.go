package domain

import (
	"database/sql"
	"github.com/MohammedKamle/banking/errs"
	"github.com/MohammedKamle/banking/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	var err error
	customers := make([]Customer, 0)
	if status == "" {
		findAllSql := "select customer_id, name, date_of_birth, city, " +
			"zipcode, status from customers"
		// Select method of sqlx will run the query and store all rows of customers table automatically
		// in customers map, in other words we don't need to individually scan all the rows and append
		// each of them one by one as we did earlier
		err = d.client.Select(&customers, findAllSql)
	} else {
		findAllSql := "select customer_id, name, date_of_birth, city, zipcode, status" +
			" from customers where status = ?"
		err = d.client.Select(&customers, findAllSql, status)
	}
	if err != nil {
		logger.Error("Error while querying customer table " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}
	return customers, nil
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	customerSql := "select customer_id, name, date_of_birth, city, zipcode, " +
		"status from customers where customer_id = ?"
	var c Customer
	err := d.client.Get(&c, customerSql, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("customer not found")
		} else {
			logger.Error("Error while scanning customer " + err.Error())
			return nil, errs.NewUnexpectedError("unexpected database error")
		}
	}
	return &c, nil
}

func NewCustomerRepositoryDb(dbClient *sqlx.DB) CustomerRepositoryDb {
	return CustomerRepositoryDb{dbClient}
}
