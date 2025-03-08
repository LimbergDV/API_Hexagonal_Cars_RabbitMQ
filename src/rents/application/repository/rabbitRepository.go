package repository

import "api-hexagonal-cars/src/rents/domain"

type IRabbit interface {
	RentCar(rent domain.Rent)
	ReturnRent(rent domain.Rent)
}