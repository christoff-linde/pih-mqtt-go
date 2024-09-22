-- name: CreateSensorReading :one
INSERT INTO sensor_readings (reading_timestamp, sensor_id, temperature, humidity, pressure)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;
