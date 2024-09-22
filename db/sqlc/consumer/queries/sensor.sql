-- name: CreateSensor :one
INSERT INTO sensors (sensor_name, sensor_unique_id, sensor_location, sensor_type, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

---- name: GetSensors :many
--SELECT * FROM sensors LIMIT $1 OFFSET $2;
--
---- name: GetSensorById :one
--SELECT * FROM sensors WHERE id=$1;

