package application

import "api-hexagonal-cars/src/rents/domain"



type GetAllRents struct{
	db domain.IRent
}

func NewGetAllRents(db domain.IRent) *GetAllRents {
	return &GetAllRents{db: db}
}

func (gar *GetAllRents) Run() []domain.Rent {
	return gar.db.GetAllRents()
}