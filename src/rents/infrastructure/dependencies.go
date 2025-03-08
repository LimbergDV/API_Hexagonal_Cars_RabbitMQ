package infrastructure

import "api-hexagonal-cars/src/rents/infrastructure/adapters"



var mysql *MySQL
var rabbitmq *adapters.RabbitMQ

func GoDependences() {
	mysql = NewMySQL()
	rabbitmq = adapters.NewRabbitMQ()
}

func GetMySQL() *MySQL {
	return mysql
}

func GetRabbitMQ() *adapters.RabbitMQ {
	return rabbitmq
}