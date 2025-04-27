package models

type Ticket struct {
	TID               uint32        `json:"tid"`
	ServiceID         string        `json:"service_id"`
	SeatNo            uint32        `json:"seat"`
	ServiceName       string        `json:"service_name"`
	ServiceProviderID string        `json:"service_provider"`
    Status            string        `json:"status"`
	PassengerID       string        `json:"passenger"`
	Price             uint32        `json:"price"`
	StartTime         string        `json:"start_time"`
	Duration          uint32        `json:"duration"`
	Source            string        `json:"source"`
	Destination       string        `json:"destination"`
	TransportType     string        `json:"transport_type"`
}
