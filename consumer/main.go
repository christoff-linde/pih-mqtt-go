package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	db "github.com/christoff-linde/pih-core-go/consumer/database"
	"github.com/jackc/pgx/v5/pgtype"
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

func (appConfig *appConfig) handleGetSensorByName(name string) (db.Sensor, error) {
	sensor, err := appConfig.DB.GetSensorByName(context.Background(), name)

	return sensor, err
}

func (appConfig *appConfig) handleCreateSensorMetadata(sensor db.Sensor) db.SensorMetadatum {
	metadata, err := appConfig.DB.CreateSensorMetadata(context.Background(), db.CreateSensorMetadataParams{
		ID:       1,
		SensorID: sensor.ID,
	})
	failOnError(err, "Failed to create sensorMetadata")

	return metadata
}

type DeviceData struct {
	DeviceID    string  `json:"device_id"`
	Temperature float64 `json:"temperature"`
	Humidity    float64 `json:"humidity"`
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
	// sensorMetadata := appCfg.handleCreateSensorMetadata(sensor)
	// fmt.Println("Created sensorMetadata:", sensorMetadata)

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
	temperatureQueue, err := channel.QueueDeclare("temperature", true, false, false, false, nil)
	failOnError(err, "Faild to declare temperatureQueue")

	// Delcare humidityQueue
	humidityQueue, err := channel.QueueDeclare("humidity", true, false, false, false, nil)
	failOnError(err, "Faild to declare humidityQueue")

	// iotQueue
	iotQueue, err := channel.QueueDeclare("iot", true, false, false, false, nil)
	failOnError(err, "Failed to declare iot queue")

	err = channel.QueueBind(iotQueue.Name, "pih", "iot", false, nil)
	err = channel.QueueBind(temperatureQueue.Name, "pih.temperature", "iot", false, nil)
	err = channel.QueueBind(humidityQueue.Name, "pih.humidity", "iot", false, nil)

	// Get temperatureQueue msgs
	// temperatureMsgs, err := channel.Consume(temperatureQueue.Name, "", true, false, false, false, nil)
	// Get humidityQueue msgs
	// humidityMsgs, err := channel.Consume(humidityQueue.Name, "", true, false, false, false, nil)
	// Get all iot msgs
	iotMsgs, err := channel.ConsumeWithContext(context.Background(), iotQueue.Name, "", true, false, false, false, nil)

	// Forever loop
	var forever chan struct{}

	go func() {
		for d := range iotMsgs {
			log.Printf("Received a message from: %v: %v", iotQueue.Name, d.Body)

			var deviceData DeviceData
			err := json.Unmarshal([]byte(d.Body), &deviceData)
			if err != nil {
				log.Printf("Error parsing JSON: %v", err)
			}

			sensor, err := appCfg.handleGetSensorByName(deviceData.DeviceID)
			if err != nil {
				log.Printf("Sensor %v not found: %v", deviceData.DeviceID, err)
			} else {
				sensorReading, err := appCfg.DB.CreateSensorReading(context.Background(), db.CreateSensorReadingParams{
					SensorID: pgtype.Int4{
						Int32: sensor.ID,
						Valid: true,
					},
					Temperature: pgtype.Float8{
						Float64: deviceData.Temperature,
						Valid:   true,
					},
					Humidity: pgtype.Float8{
						Float64: deviceData.Humidity,
						Valid:   true,
					},
					Pressure: pgtype.Float8{},
				})
				if err != nil {
					log.Printf("Failed to create sensorReading: %v", err)
				}
				fmt.Printf("Added: %v", sensorReading)
			}
		}
	}()

	//// Consume temperatureMsgs
	//go func() {
	//	for d := range temperatureMsgs {
	//		log.Printf("Received a message from %s: %s", temperatureQueue.Name, d.Body)

	//		var deviceData DeviceData
	//		err := json.Unmarshal([]byte(d.Body), &deviceData)
	//		if err != nil {
	//			log.Printf("Error parsing JSON: %v", err)
	//		}

	//		sensorReading, err := appCfg.handleCreateSensorReading(&sensor, &db.SensorReading{
	//			ReadingTimestamp: pgtype.Timestamptz{
	//				Time:             d.Timestamp,
	//				InfinityModifier: 0,
	//				Valid:            true,
	//			},
	//			SensorID: pgtype.Int4{
	//				Int32: sensor.ID,
	//				Valid: true,
	//			},
	//			Temperature: pgtype.Float8{
	//				Float64: deviceData.Data[0],
	//				Valid:   true,
	//			},
	//		})
	//		if err != nil {
	//			log.Printf("Failed to create sensorReading: %v", err)
	//		}
	//		fmt.Printf("Added: %v", sensorReading)
	//	}
	//}()

	//// Consume humidityMsgs
	//go func() {
	//	for d := range humidityMsgs {
	//		log.Printf("Received a message from %s: %s", humidityQueue.Name, d.Body)

	//		var deviceData DeviceData
	//		err := json.Unmarshal([]byte(d.Body), &deviceData)
	//		if err != nil {
	//			log.Printf("Error parsing JSON: %v", err)
	//		}

	//		sensorReading, err := appCfg.handleCreateSensorReading(&sensor, &db.SensorReading{
	//			ReadingTimestamp: pgtype.Timestamptz{
	//				Time:             d.Timestamp,
	//				InfinityModifier: 0,
	//				Valid:            true,
	//			},
	//			SensorID: pgtype.Int4{
	//				Int32: sensor.ID,
	//				Valid: true,
	//			},
	//			Humidity: pgtype.Float8{
	//				Float64: deviceData.Data[0],
	//				Valid:   true,
	//			},
	//		})
	//		if err != nil {
	//			log.Printf("Failed to create sensorReading: %v", err)
	//		}
	//		fmt.Printf("Added: %v", sensorReading)
	//	}
	//}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C.")
	<-forever
}
