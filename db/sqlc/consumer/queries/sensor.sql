-- name: CreateSensor :one
INSERT INTO sensors (sensor_name, created_at, updated_at)
VALUES ($1, $2, $3)
RETURNING *;

---- name: GetSensors :many
--SELECT * FROM sensors LIMIT $1 OFFSET $2;
--
-- name: GetSensorById :one
SELECT * FROM sensors WHERE id=$1;

