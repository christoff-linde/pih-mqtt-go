-- name: CreateSensorMetadata :one
INSERT INTO sensor_metadata ( id, sensor_id, manufacturer, model_number, installation_time, updated_at, additional_data )
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;
