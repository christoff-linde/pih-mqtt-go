-- name: CreateSensor :one
INSERT INTO sensors (sensor_name, sensor_location, sensor_type)
VALUES ($1, $2, $3)
RETURNING *;

---- name: GetSensors :many
--SELECT * FROM sensors LIMIT $1 OFFSET $2;
--
-- name: GetSensorById :one
SELECT * FROM sensors WHERE id=$1;

