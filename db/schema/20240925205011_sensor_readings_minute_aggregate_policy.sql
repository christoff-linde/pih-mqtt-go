-- +goose NO TRANSACTION
-- +goose Up
-- +goose StatementBegin
SELECT add_continuous_aggregate_policy('sensor_readings_minutes',
  start_offset => INTERVAL '1 month',
  end_offset => INTERVAL '1 minute',
  schedule_interval => INTERVAL '1 minute');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT remove_continuous_aggregate_policy(
    'sensor_readings_minutes'
);
-- +goose StatementEnd
