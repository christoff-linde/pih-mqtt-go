package server

import (
	"time"
)

type Sensor struct {
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	SensorName string    `json:"sensor_name"`
	ID         int32     `json:"id,omitempty"`
}
