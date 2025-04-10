package controllers

import (
	"api-hexagonal-cars/src/rents/application/services"
	application "api-hexagonal-cars/src/rents/application/useCases"
	"api-hexagonal-cars/src/rents/domain"
	"api-hexagonal-cars/src/rents/infrastructure"
	"api-hexagonal-cars/src/rents/infrastructure/controllers/validators"
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateRentController struct {
	app     *application.CreateRent
	service *services.RentCarEvent
}

func NewCreateRentController() *CreateRentController {
	mysql := infrastructure.GetMySQL()
	app := application.NewCreateRent(mysql)
	rabbit := infrastructure.GetRabbitMQ()
	service := services.NewRentCarEvent(rabbit)
	return &CreateRentController{app: app, service: service}
}

func (cr_c *CreateRentController) AddRent(c *gin.Context) {
	var rent domain.Rent

	if err := c.ShouldBindJSON(&rent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"error":  "Datos inválidos: " + err.Error(),
		})
		return
	}

	if err := validators.VerifyRent(rent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"error":  "Datos inválidos: " + err.Error(),
		})
		return
	}

	fmt.Println(rent.Show())

	_, err := cr_c.app.Run(rent)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"error":  "No se pudo guardar la renta " + err.Error(),
		})
		return
	}

	// Mandar al RentProcessor
	cr_c.service.Run(rent)

	c.JSON(http.StatusCreated, gin.H {"mensaje": "Renta creada"})
	c.JSON(http.StatusOK, rent)

}
