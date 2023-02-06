package domain

import (
	"github.com/MohammedKamle/banking/dto"
	"github.com/MohammedKamle/banking/errs"
)

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

func (c Customer) CovertToDto() dto.CustomerResponse {
	return dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateOfBirth: c.DateOfBirth,
		Status:      c.statusAsText(),
	}
}

func (c Customer) statusAsText() string {
	// replacing status as active/inactive from 1/0 to send in the response
	statusAsText := "active"
	if c.Status == "0" {
		statusAsText = "inactive"
	}
	return statusAsText
}
