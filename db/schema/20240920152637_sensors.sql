-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS sensors (
    id SERIAL NOT NULL PRIMARY KEY,
    sensor_name TEXT UNIQUE NOT NULL,
    sensor_location TEXT,
    sensor_type TEXT,
    created_at TIMESTAMPTZ DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT NOW() NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS sensors;
-- +goose StatementEnd
