package models

import (
	"time"
)

type Ticket struct {
	TID               uint32        `json:"tid"`
	ServiceID         string        `json:"service_id"`
	SeatNo            uint32        `json:"seat"`
	ServiceName       string        `json:"service_name"`
	ServiceProviderID string        `json:"service_provider"`
    Status            string        `json:"status"`
	PassengerID       string        `json:"passenger"`
	Price             uint32        `json:"price"`
	StartTime         time.Time     `json:"start_time"`
	Duration          time.Duration `json:"duration"`
	Source            string        `json:"source"`
	Destination       string        `json:"destination"`
	TransportType     string        `json:"transport_type"`
}
