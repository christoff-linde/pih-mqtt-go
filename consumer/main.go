package main

import (
	"context"
	"fmt"
	"log"
	"os"

	db "github.com/christoff-linde/pih-core-go/consumer/database"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
)

type appConfig struct {
	DB *db.Queries
}

func initDb(databaseUrl string) *db.Queries {
	connection, err := pgxpool.New(context.Background(), databaseUrl)
	if err != nil {
		log.Panic(err, "Unable to connect to database: %v\n", err)
	}

	db := db.New(connection)
	return db
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func (appConfig *appConfig) handleCreateSensor() (db.Sensor, error) {
	sensor, err := appConfig.DB.CreateSensor(context.Background(), db.CreateSensorParams{
		SensorName: "esp32-test-02",
	})

	return sensor, err
}

func (appConfig *appConfig) handleGetSensorBySensorId(id int32) db.Sensor {
	sensor, err := appConfig.DB.GetSensorById(context.Background(), id)
	failOnError(err, "Could not fetch sensor")

	return sensor
}

//CREATE TABLE IF NOT EXISTS sensor_metadata (
//    id INT NOT NULL,
//    sensor_id INT NOT NULL,
//    manufacturer TEXT,
//    model_number TEXT,
//    installation_time TIMESTAMPTZ NOT NULL DEFAULT NOW(),
//    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
//    additional_data JSONB
//);

func (appConfig *appConfig) handleCreateSensorMetadata(sensor db.Sensor) db.SensorMetadatum {
	metadata, err := appConfig.DB.CreateSensorMetadata(context.Background(), db.CreateSensorMetadataParams{
		ID:       1,
		SensorID: sensor.ID,
	})
	failOnError(err, "Failed to create sensorMetadata")

	return metadata
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	brokerUrl := os.Getenv("BROKER_URL")
	fmt.Println("Found BROKER_URL with value:", brokerUrl)

	databaseUrl := os.Getenv("DB_URL")
	fmt.Println("Found DB_URL with value:", databaseUrl)

	dbConn := initDb(databaseUrl)
	appCfg := appConfig{DB: dbConn}

	sensor, err := appCfg.handleCreateSensor()
	if err != nil {
		sensor = appCfg.handleGetSensorBySensorId(1)
		fmt.Println("Fetched sensor:", sensor)
	} else {
		fmt.Println("Created sensor:", sensor)
	}
	sensorMetadata := appCfg.handleCreateSensorMetadata(sensor)
	fmt.Println("Created sensorMetadata:", sensorMetadata)

	return

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
