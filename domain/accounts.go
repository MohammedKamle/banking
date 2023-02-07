package domain

import (
	"github.com/MohammedKamle/banking/dto"
	"github.com/MohammedKamle/banking/errs"
)

type Account struct {
	AccountId   string
	CustomerId  string
	OpeningDate string
	AccountType string
	Amount      float64
	Status      string
}

type AccountRepository interface {
	Save(account Account) (*Account, *errs.AppError)
}

func (a Account) ToNewAccountResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{AccountID: a.AccountId}
}
