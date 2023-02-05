package domain

import "github.com/MohammedKamle/banking/errs"

// db is used as our struct has key Id and database has customer_id, so sqlx will throw error while scanning for
// Id and DateOfBirth
type Customer struct {
	Id          string `db:"customer_id"`
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string `db:"date_of_birth"`
	Status      string
}

type CustomerRepository interface {
	FindAll(status string) ([]Customer, *errs.AppError)
	// ById function will return a pointer to customer because we want to send nil if
	//no Customer is available against the id provided  and that is possible only with pointer
	ById(id string) (*Customer, *errs.AppError)
}
