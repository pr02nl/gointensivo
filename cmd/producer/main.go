package main

import (
	"encoding/json"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/pr02nl/gointensivo/internal/order/entity"
	amqp "github.com/rabbitmq/amqp091-go"
)

func Publish(ch *amqp.Channel, order entity.Order) error {
	body, err := json.Marshal(order)
	if err != nil {
		return err
	}
	err = ch.Publish(
		"amq.direct", // exchange
		"",           // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		}, // message
	)
	if err != nil {
		return err
	}
	return nil
}

func GenerateOrders() entity.Order {
	return entity.Order{
		ID:    uuid.New().String(),
		Price: rand.Float64() * 100,
		Tax:   rand.Float64() * 10,
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()
	for i := 0; i < 10000000; i++ {
		order := GenerateOrders()
		err := Publish(ch, order)
		if err != nil {
			panic(err)
		}
		time.Sleep(300 * time.Millisecond)
	}
}
