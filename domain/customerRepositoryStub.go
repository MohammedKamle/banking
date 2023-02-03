package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1001", "Ashish", "New Delhi", "2654365", "2000-01-01", "1"},
		{"1002", "Sonu", "Mumbai", "672284", "1996-01-01", "1"},
		{"1003", "Daniyal", "Banglore", "765456", "1997-01-01", "1"},
	}
	return CustomerRepositoryStub{customers}
}
