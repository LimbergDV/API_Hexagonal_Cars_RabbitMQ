package application

import "api-hexagonal-cars/src/rents/domain"

type CreateRent struct{
	db domain.IRent
}

func NewCreateRent(db domain.IRent) *CreateRent {
	return &CreateRent{db: db}
}

func (cr CreateRent) Run(rent domain.Rent) (uint, error) {
	return cr.db.CreateRent(rent)
}