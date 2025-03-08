package domain

type IRent interface {
	CreateRent(rent Rent) (uint, error)
	GetAllRents() []Rent
	DeleteRent(id_customer int, id_car int) (uint, error)
}