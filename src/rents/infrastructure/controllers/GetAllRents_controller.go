package controllers

import (
	application "api-hexagonal-cars/src/rents/application/useCases"
	"api-hexagonal-cars/src/rents/infrastructure"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllRentsController struct {
	app *application.GetAllRents
}

func NewGetAllRentsController() *GetAllRentsController {
	mysql := infrastructure.GetMySQL()
	app := application.NewGetAllRents(mysql)
	return &GetAllRentsController{app: app}
}

func (galr *GetAllRentsController) GetAllRents(c *gin.Context) {
	res := galr.app.Run()

	if res == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": false,
			"error": "No se consigió resultados",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"rentas de carros": res})
}