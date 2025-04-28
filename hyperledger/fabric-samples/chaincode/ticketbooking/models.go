package main


   type User struct {
       DocType string  `json:"docType"`
       Email   string  `json:"email"`
       Name    string  `json:"name"`
       Role    string  `json:"role"`
       Balance float64 `json:"balance"`
   }

   type Ticket struct {
       DocType           string        `json:"docType"`
       TID               uint32        `json:"tid"`
       ServiceID         string        `json:"service_id"`
       SeatNo            uint32        `json:"seat_no"`
       ServiceName       string        `json:"service_name"`
       ServiceProviderID string        `json:"service_provider"`
       Status            string        `json:"status"`
       PassengerID       string        `json:"passenger"`
       Price             float64       `json:"price"`
       StartTime         string     `json:"start_time"`
       Duration          uint32 `json:"duration"`
       Source            string        `json:"source"`
       Destination       string        `json:"destination"`
       TransportType     string        `json:"transport_type"`
   }