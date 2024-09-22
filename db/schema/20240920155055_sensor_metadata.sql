-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS sensor_metadata (
    id SERIAL NOT NULL PRIMARY KEY,
    sensor_id INT NOT NULL UNIQUE,
    manufacturer TEXT,
    model_number TEXT,
    installation_time TIMESTAMPTZ DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT NOW() NOT NULL,
    additional_data JSONB
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS sensor_metadata;
-- +goose StatementEnd
