package application

import "api-hexagonal-cars/src/customers/domain"



type GetAllCustomers struct {
	db domain.ICustomer
}

func NewGetAllCustomers(db domain.ICustomer) *GetAllCustomers {
	return &GetAllCustomers{db: db}
}

func (lc *GetAllCustomers) Run () []domain.Customer {
	return lc.db.GetAll()
}