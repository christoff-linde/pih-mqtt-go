-- name: CreateSensor :one
INSERT INTO sensors (sensor_name, created_at, updated_at)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetSensors :many
SELECT * FROM sensors LIMIT $1 OFFSET $2;

-- name: GetSensorById :one
SELECT * FROM sensors WHERE id=$1;

-- name: GetSensorByName :one
SELECT * FROM sensors WHERE sensor_name=$1;

-- name: UpdateSensor :execresult
UPDATE sensors
SET sensor_name = $2, updated_at = now()
WHERE id = $1
RETURNING *;

-- name: DeleteSensor :execresult
DELETE FROM sensors WHERE id=$1;
