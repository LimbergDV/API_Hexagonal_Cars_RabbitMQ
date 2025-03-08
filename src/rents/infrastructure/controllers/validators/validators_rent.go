package validators

import (
	"api-hexagonal-cars/src/rents/domain"
	"errors"
)

func VerifyRent(rent domain.Rent) error {
	if rent.Id_Car <= 0 {
		return errors.New("id_car invalidated")
	}
	if rent.Return_date_rent == "" {
		return errors.New("verify the date of return date")
	}
	if rent.Id_Customer <= 0 {
		return errors.New("id_customer invalidated")
	}
	return nil
}