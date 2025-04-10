package adapters

import (
	"api-hexagonal-cars/src/rents/domain"
	"context"
	"encoding/json"

	"log"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	conn *amqp.Connection
	ch *amqp.Channel
}

func NewRabbitMQ() *RabbitMQ {
	conn, err := amqp.Dial(os.Getenv("URL_RABBIT"))
	failOnError(err, "Failed to connect to RabbitMQ")
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

	return &RabbitMQ{conn: conn, ch: ch} 
}

func (r *RabbitMQ) RentCar(rent domain.Rent) {
	payload, err := json.Marshal(rent)
	failOnError(err, "Error al serializar Renta de carros a JSON")
	r.prepareToMessage(payload)
}

func (r *RabbitMQ) ReturnCar(rent domain.Rent) {
	payload, err := json.Marshal(rent)
	failOnError(err, "Error al serializar Renta de carros a JSON")
	r.prepareToMessage(payload)
}

func (r *RabbitMQ) prepareToMessage(body []byte) {
	// Declaración del exchange (intercambiador):
	r.ch.ExchangeDeclare(
		"exchanges_cars",   // name
		"topic", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	  
	r.ch.PublishWithContext(ctx,
		"exchanges_cars",     // exchange
		"rents", // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing {
		  ContentType: "application/json",
		  Body:        body,
		})
	log.Printf(" [x] Sent %s\n", body)
}

func failOnError(err error, msg string) {
	if err != nil {
	  log.Panicf("%s: %s", msg, err)
	}
}