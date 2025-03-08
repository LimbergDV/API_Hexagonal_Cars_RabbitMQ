package main

import (
	customers "api-hexagonal-cars/src/customers/infrastructure"
	routesCustomers "api-hexagonal-cars/src/customers/infrastructure/routes"
	infraCustomers "api-hexagonal-cars/src/rents/infrastructure"
	routesRents "api-hexagonal-cars/src/rents/infrastructure/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main () {
	infraCustomers.GoDependences()
	customers.GoMySQL()

	r := gin.Default()
	r.Use((cors.Default()))

	routesCustomers.Routes(r)
	routesRents.Routes(r)
	
	r.Run(":8080")
}