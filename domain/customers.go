package domain

import "github.com/MohammedKamle/banking/errs"

type Customer struct {
	Id          string
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string
	Status      string
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
	// ById function will return a pointer to customer because we want to send nil if
	//no Customer is available against the id provided  and that is possible only with pointer
	ById(id string) (*Customer, *errs.AppError)
}
