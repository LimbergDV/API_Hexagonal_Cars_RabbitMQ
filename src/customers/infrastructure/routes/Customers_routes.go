package routes

import (
	"api-hexagonal-cars/src/customers/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

func Routes (r *gin.Engine) {
	
	customersRoutes := r.Group("/customers") 
	{
		customersRoutes.POST("/", controllers.NewCreateCustomerController().Run)
		customersRoutes.GET("/", controllers.NewGetAllCustomersController().Run)
		customersRoutes.PUT("/:id", controllers.NewUpdateCustomerByIdController().Run)
		customersRoutes.DELETE("/:id", controllers.NewDeleteCustomerByIdController().Run)
	}
}