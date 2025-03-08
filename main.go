package main

import (
	
	routesCustomers "api-hexagonal-cars/src/customers/infrastructure/routes"
	customers "api-hexagonal-cars/src/customers/infrastructure"
	
	

	"github.com/gin-gonic/gin"
)

func main () {
	customers.GoMySQL()

	r := gin.Default()

	routesCustomers.Routes(r)
	
	r.Run(":8080")
}