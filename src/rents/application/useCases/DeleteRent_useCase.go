package application

import "api-hexagonal-cars/src/rents/domain"

type DeleteRent struct {
	db domain.IRent
}

func NewDeleteRent(db domain.IRent) *DeleteRent {
	return &DeleteRent{db: db}
}

func (dr *DeleteRent) Run(id_customer int, id_car int) (uint, error) {
	return dr.db.DeleteRent(id_customer, id_car)
}