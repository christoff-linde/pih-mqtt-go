-- +goose NO TRANSACTION
-- +goose Up
-- +goose StatementBegin
SELECT add_continuous_aggregate_policy('sensor_readings_hourly',
  start_offset => INTERVAL '1 month',
  end_offset => INTERVAL '1 hour',
  schedule_interval => INTERVAL '15 minute');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT remove_continuous_aggregate_policy(
    'sensor_readings_hourly'
);
-- +goose StatementEnd
