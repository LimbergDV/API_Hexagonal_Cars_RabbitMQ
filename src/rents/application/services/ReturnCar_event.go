package services

import (
	"api-hexagonal-cars/src/rents/application/repository"
	"api-hexagonal-cars/src/rents/domain"
)

type ReturnCarEvent struct {
	rmq repository.IRabbit
}

func NewReturnRentEvent(rmq repository.IRabbit) *ReturnCarEvent {
	return &ReturnCarEvent{rmq: rmq}
}

func (rce *ReturnCarEvent) Run(rent domain.Rent) {
	rce.rmq.ReturnRent(rent)
}