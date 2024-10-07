package server

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type Sensor struct {
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	SensorName string    `json:"sensor_name"`
	ID         int32     `json:"id,omitempty"`
}

type SensorMetadata struct {
	ID             int32       `json:"id"`
	SensorID       int32       `json:"sensor_id"`
	SensorType     pgtype.Text `json:"sensor_type"`
	Manufacturer   pgtype.Text `json:"manufacturer"`
	ModelNumber    pgtype.Text `json:"model_number"`
	SensorLocation pgtype.Text `json:"sensor_location"`
	AdditionalData []byte      `json:"additional_data"`
}
