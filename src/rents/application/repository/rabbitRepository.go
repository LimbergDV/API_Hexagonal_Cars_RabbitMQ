package repository

import "api-hexagonal-cars/src/rents/domain"

type IRabbit interface {
	RentCar(rent domain.Rent)
	ReturnCar(rent domain.Rent)
}