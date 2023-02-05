package service

import (
	"github.com/MohammedKamle/banking/domain"
	"github.com/MohammedKamle/banking/errs"
)

type CustomerService interface {
	GetAllCustomers(status string) ([]domain.Customer, *errs.AppError)
	GetCustomer(id string) (*domain.Customer, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers(status string) ([]domain.Customer, *errs.AppError) {
	// client will be sending status as active/inactive, but we are storing them as 1 or 0,
	// so we need to convert them appropriately
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}
	return s.repo.FindAll(status)
}

func (s DefaultCustomerService) GetCustomer(id string) (*domain.Customer, *errs.AppError) {
	return s.repo.ById(id)
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
