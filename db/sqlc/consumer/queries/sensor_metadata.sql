-- name: CreateSensorMetadata :one
INSERT INTO sensor_metadata ( id, sensor_id, manufacturer, model_number,  additional_data )
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetSensorMetadataForSensorId :one
SELECT sensors.*
FROM sensors
         JOIN sensor_metadata ON sensors.id = sensor_metadata.sensor_id
WHERE sensor_metadata.sensor_id = $1
LIMIT $2;
