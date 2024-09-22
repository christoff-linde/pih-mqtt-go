-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS sensor_metadata (
    id SERIAL NOT NULL PRIMARY KEY,
    manufacturer TEXT,
    model_number TEXT,
    installation_time TIMESTAMPTZ DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT NOW() NOT NULL,
    additional_data JSONB,
    sensor_id SERIAL NOT NULL REFERENCES sensors(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS sensor_metadata;
-- +goose StatementEnd
