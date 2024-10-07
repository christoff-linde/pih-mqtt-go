-- name: CreateSensorMetadata :one
INSERT INTO sensor_metadata ( id, sensor_id, sensor_type, manufacturer, model_number, sensor_location, additional_data )
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: GetSensorMetadata :many
SELECT *
FROM sensor_metadata
LIMIT $1 OFFSET $2;

-- name: GetSensorMetadataForSensorId :one
SELECT sensors.*
FROM sensors
         JOIN sensor_metadata ON sensors.id = sensor_metadata.sensor_id
WHERE sensor_metadata.sensor_id = $1;

-- name: UpdateSensorMetadata :execresult
UPDATE sensor_metadata
SET sensor_id = $1, sensor_type = $3, manufacturer = $4, model_number = $5, sensor_location = $6, additional_data = $7
WHERE id = $2
RETURNING *;

-- name: DeleteSensorMetadata :execresult
DELETE FROM sensor_metadata WHERE id = $1;


-- -- name: CreateSensor :one
-- INSERT INTO sensors (sensor_name, created_at, updated_at)
-- VALUES ($1, $2, $3)
-- RETURNING *;
-- 
-- -- name: GetSensors :many
-- SELECT * FROM sensors LIMIT $1 OFFSET $2;
-- 
-- -- name: GetSensorById :one
-- SELECT * FROM sensors WHERE id=$1;
-- 
-- -- name: GetSensorByName :one
-- SELECT * FROM sensors WHERE sensor_name=$1;
-- 
-- -- name: UpdateSensor :execresult
-- UPDATE sensors
-- SET sensor_name = $2, updated_at = now()
-- WHERE id = $1
-- RETURNING *;
-- 
-- -- name: DeleteSensor :execresult
-- DELETE FROM sensors WHERE id=$1;
