-- name: CreateSensorReading :one
INSERT INTO sensor_readings (sensor_id, temperature, humidity, pressure)
VALUES ($1, $2, $3, $4 )
RETURNING *;
