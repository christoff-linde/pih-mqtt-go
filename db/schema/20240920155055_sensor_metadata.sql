-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS sensor_metadata (
    id INT NOT NULL,
    sensor_id INT NOT NULL,
    manufacturer TEXT,
    model_number TEXT,
    installation_time TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    additional_data JSONB
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS sensor_metadata;
-- +goose StatementEnd
