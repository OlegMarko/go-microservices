package main

import (
	"fmt"
	"listener/event"
	"log"
	"math"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	rabbitConn, err := connect()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	defer rabbitConn.Close()

	log.Println("Listening and consuming RabbitMQ messages")

	consumer, err := event.NewConsumer(rabbitConn)
	if err != nil {
		panic(err)
	}

	err = consumer.Listen([]string{"log.INFO", "log.WARNING", "log.ERROR"})
	if err != nil {
		log.Println(err)
	}
}

func connect() (*amqp.Connection, error) {
	var counts int64
	var backOff = 1 * time.Second
	var connection *amqp.Connection

	rabbitMQURL := "amqp://guest:guest@rabbitmq"

	for {
		c, err := amqp.Dial(rabbitMQURL)
		if err != nil {
			fmt.Println("RabbitMQ not yeat ready...")
			log.Println(err)

			counts++
		} else {
			log.Println("Connected to RabbitMQ")

			connection = c

			break
		}

		if counts > 5 {
			fmt.Println(err)
			return nil, err
		}

		backOff = time.Duration(math.Pow(float64(counts), 2)) * time.Second
		log.Printf("backing of %d ...", backOff)
		time.Sleep(backOff)

		continue
	}

	return connection, nil
}
