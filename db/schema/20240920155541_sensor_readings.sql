-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS sensor_readings (
    reading_timestamp TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    sensor_id INTEGER,
    temperature DOUBLE PRECISION,
    humidity DOUBLE PRECISION,
    pressure DOUBLE PRECISION
);

SELECT CREATE_HYPERTABLE('sensor_readings', 'reading_timestamp');

CREATE UNIQUE INDEX sensor_readings_sensor_id_reading_timestamp_idx ON sensor_readings (sensor_id, reading_timestamp DESC);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS sensor_readings;
-- +goose StatementEnd
