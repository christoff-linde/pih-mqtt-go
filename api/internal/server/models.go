package server

import "github.com/jackc/pgx/v5/pgtype"

type Sensor struct {
	ID         int32              `json:"id"`
	SensorName string             `json:"sensor_name"`
	CreatedAt  pgtype.Timestamptz `json:"created_at"`
	UpdatedAt  pgtype.Timestamptz `json:"updated_at"`
}
