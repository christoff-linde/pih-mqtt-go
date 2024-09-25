-- +goose NO TRANSACTION
-- +goose Up
-- +goose StatementBegin
SELECT add_continuous_aggregate_policy('sensor_readings_daily',
  start_offset => INTERVAL '12 month',
  end_offset => INTERVAL '1 hour',
  schedule_interval => INTERVAL '1 hour');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT remove_continuous_aggregate_policy(
    'sensor_readings_daily'
);
-- +goose StatementEnd
