package routes

import (
	"api-hexagonal-cars/src/rents/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	rentsRouter := r.Group("/rents") 
	{
		rentsRouter.POST("/", controllers.NewCreateRentController().AddRent)
		rentsRouter.GET("/", controllers.NewGetAllRentsController().GetAllRents)
		rentsRouter.DELETE("/", controllers.NewDeleteRentController().DeleteRent)
	}
}