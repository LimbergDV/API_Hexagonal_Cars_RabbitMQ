package controllers

import (
	"api-hexagonal-cars/src/rents/application/services"
	application "api-hexagonal-cars/src/rents/application/useCases"
	"api-hexagonal-cars/src/rents/domain"
	"api-hexagonal-cars/src/rents/infrastructure"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteRentController struct {
	app *application.DeleteRent
	service *services.ReturnCarEvent
}

func NewDeleteRentController() *DeleteRentController {
	mysql := infrastructure.GetMySQL()
	rabbit := infrastructure.GetRabbitMQ()
	app := application.NewDeleteRent(mysql)
	service := services.NewReturnRentEvent(rabbit)
	return &DeleteRentController{app: app, service: service}
}

func (drc *DeleteRentController) DeleteRent(c *gin.Context) {
	customer := c.Query("id_customer")
    car := c.Query("id_car") 

	id_customer, _ := strconv.ParseInt(customer, 10, 64)
	id_car, _ := strconv.ParseInt(car, 10, 64)

	rowsAffected, _ := drc.app.Run(int(id_customer),int(id_car))

	if rowsAffected == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"error": "No se pudo eliminar la renta: No se entontró la referencia o ocurrió algo más",
		})
		return
	}

	var rent domain.Rent
	rent.Id_Car = id_car
	rent.Id_Customer = id_customer
	rent.Return_date_rent = "0000-00-00"
	// Para crear el evento 
	drc.service.Run(rent )

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"message": "Renta eliminada",
	})
}