package domain

import (
	
	"fmt"
)

type Rent struct {
	Id_Car int64
	Id_Customer int64
	Return_date_rent string
}

func (r *Rent) Show() string {
	return fmt.Sprintf("{id_customer: %d, id_car: %d, return_date: %s}",
		r.Id_Customer, r.Id_Car, r.Return_date_rent)
}
