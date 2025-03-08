package services

import (
	"api-hexagonal-cars/src/rents/application/repository"
	"api-hexagonal-cars/src/rents/domain"
)

type RentCarEvent struct {
	rmq repository.IRabbit
}

func NewRentCarEvent(rmq repository.IRabbit) *RentCarEvent {
	return &RentCarEvent{rmq: rmq}
}

func (rce *RentCarEvent) Run(rent domain.Rent) {
	rce.rmq.RentCar(rent)
}