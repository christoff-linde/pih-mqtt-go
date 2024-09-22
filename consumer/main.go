package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	brokerUrl := os.Getenv("BROKER_URL")
	fmt.Println("Found BROKER_URL with value:", brokerUrl)

	// RabbitMQ Setup
	conn, err := amqp.Dial(brokerUrl)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	channel, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer channel.Close()

	//// Declare Exchange that reads from iot exchange
	//err = channel.ExchangeDeclare("iot_clone", "topic", true, false, false, false, nil)
	//failOnError(err, "Failed to declare exchange iot_clone")

	// Check if IOT queue exists
	err = channel.ExchangeDeclarePassive("iot", "topic", true, true, false, false, nil)
	failOnError(err, "IOT exchange does not exist")

	//// Bind to existing IOT queue
	//// This exchange is simply used for testing for now
	//err = channel.ExchangeBind("iot_clone", "#", "iot", false, nil)
	//failOnError(err, "Failed to bind to iot_clone")

	// Declare temperatureQueue
	temperatureQueue, err := channel.QueueDeclare("temperature_clone", true, false, false, false, nil)
	failOnError(err, "Faild to declare temperatureQueue")

	// Delcare humidityQueue
	humidityQueue, err := channel.QueueDeclare("humidity_clone", true, false, false, false, nil)
	failOnError(err, "Faild to declare humidityQueue")

	err = channel.QueueBind(temperatureQueue.Name, "pih.temperature", "iot", false, nil)
	err = channel.QueueBind(humidityQueue.Name, "pih.humidity", "iot", false, nil)

	// Get temperatureQueue msgs
	temperatureMsgs, err := channel.Consume(temperatureQueue.Name, "", true, false, false, false, nil)
	// Get humidityQueue msgs
	humidityMsgs, err := channel.Consume(humidityQueue.Name, "", true, false, false, false, nil)

	// Forever loop
	var forever chan struct{}

	// Consume temperatureMsgs
	go func() {
		for d := range temperatureMsgs {
			log.Printf("Received a message from %s: %s", temperatureQueue.Name, d.Body)
		}
	}()

	// Consume humidityMsgs
	go func() {
		for d := range humidityMsgs {
			log.Printf("Received a message from %s: %s", humidityQueue.Name, d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C.")
	<-forever
}
