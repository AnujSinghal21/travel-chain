package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"backend/models"
)

// GetTicketByTID fetches a ticket by TID
func GetTicketByTID(c *gin.Context) {
	tid := c.Query("tid")
	if tid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "tid is required"})
		return
	}

	// Dummy ticket response
	c.JSON(http.StatusOK, gin.H{
		"tid":          tid,
		"service_name": "Mock Transport",
		"status":       "Available",
	})
}

// GetAllTickets returns all tickets
func GetAllTickets(c *gin.Context) {
	c.JSON(http.StatusOK, []models.Ticket{
		{TID: 1, ServiceID: "S001", SeatNo: 1, ServiceName: "CityLink Express", ServiceProviderID: "P001", Status: "Available", PassengerID: "", Price: 100, Source: "CityA", Destination: "CityB", TransportType: "Bus"},
		{TID: 2, ServiceID: "S002", SeatNo: 2, ServiceName: "CityLink Express", ServiceProviderID: "P001", Status: "Booked", PassengerID: "U002", Price: 100, Source: "CityA", Destination: "CityB", TransportType: "Bus"},
		{TID: 3, ServiceID: "S003", SeatNo: 1, ServiceName: "MetroConnect", ServiceProviderID: "P002", Status: "Available", PassengerID: "", Price: 150, Source: "CityB", Destination: "CityC", TransportType: "Train"},
		{TID: 4, ServiceID: "S004", SeatNo: 2, ServiceName: "MetroConnect", ServiceProviderID: "P002", Status: "Available", PassengerID: "", Price: 150, Source: "CityB", Destination: "CityC", TransportType: "Train"},
		{TID: 5, ServiceID: "S005", SeatNo: 1, ServiceName: "SkyFly", ServiceProviderID: "P003", Status: "Available", PassengerID: "", Price: 500, Source: "CityC", Destination: "CityD", TransportType: "Flight"},
		{TID: 6, ServiceID: "S006", SeatNo: 2, ServiceName: "SkyFly", ServiceProviderID: "P003", Status: "Available", PassengerID: "", Price: 500, Source: "CityC", Destination: "CityD", TransportType: "Flight"},
		{TID: 7, ServiceID: "S007", SeatNo: 1, ServiceName: "RoadRunner", ServiceProviderID: "P004", Status: "Available", PassengerID: "", Price: 80, Source: "CityD", Destination: "CityE", TransportType: "Bus"},
		{TID: 8, ServiceID: "S008", SeatNo: 1, ServiceName: "RoadRunner", ServiceProviderID: "P004", Status: "Booked", PassengerID: "U003", Price: 80, Source: "CityD", Destination: "CityE", TransportType: "Bus"},
		{TID: 9, ServiceID: "S009", SeatNo: 1, ServiceName: "UrbanRide", ServiceProviderID: "P005", Status: "Available", PassengerID: "", Price: 50, Source: "CityE", Destination: "CityF", TransportType: "Bus"},
		{TID: 10, ServiceID: "S010", SeatNo: 2, ServiceName: "UrbanRide", ServiceProviderID: "P005", Status: "Available", PassengerID: "", Price: 50, Source: "CityE", Destination: "CityF", TransportType: "Bus"},
		{TID: 11, ServiceID: "S011", SeatNo: 1, ServiceName: "ExpressRail", ServiceProviderID: "P006", Status: "Available", PassengerID: "", Price: 200, Source: "CityF", Destination: "CityG", TransportType: "Train"},
		{TID: 12, ServiceID: "S012", SeatNo: 2, ServiceName: "ExpressRail", ServiceProviderID: "P006", Status: "Available", PassengerID: "", Price: 200, Source: "CityF", Destination: "CityG", TransportType: "Train"},
		{TID: 13, ServiceID: "S013", SeatNo: 1, ServiceName: "QuickFly", ServiceProviderID: "P007", Status: "Available", PassengerID: "", Price: 600, Source: "CityG", Destination: "CityH", TransportType: "Flight"},
		{TID: 14, ServiceID: "S014", SeatNo: 2, ServiceName: "QuickFly", ServiceProviderID: "P007", Status: "Available", PassengerID: "", Price: 600, Source: "CityG", Destination: "CityH", TransportType: "Flight"},
		{TID: 15, ServiceID: "S015", SeatNo: 1, ServiceName: "MegaBus", ServiceProviderID: "P008", Status: "Available", PassengerID: "", Price: 120, Source: "CityH", Destination: "CityA", TransportType: "Bus"},
	})
}


// CreateTickets creates tickets
func CreateTickets(c *gin.Context) {
	var tickets []models.Ticket
	if err := c.ShouldBindJSON(&tickets); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ticket list"})
		return
	}

	// Just return success
	c.JSON(http.StatusOK, gin.H{"message": "Tickets created successfully"})
}

// DeleteTickets deletes tickets
func DeleteTickets(c *gin.Context) {
	var request struct {
		TIDs []uint32 `json:"tids"`
	}
	if err := c.ShouldBindJSON(&request); err != nil || len(request.TIDs) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "TIDs are required"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tickets deleted successfully"})
}

// CancelTickets cancels tickets
func CancelTickets(c *gin.Context) {
	var request struct {
		TIDs []uint32 `json:"tids"`
	}
	if err := c.ShouldBindJSON(&request); err != nil || len(request.TIDs) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "TIDs are required"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tickets cancelled successfully"})
}

// BookTickets books tickets
func BookTickets(c *gin.Context) {
	var request struct {
		TIDs []uint32 `json:"tids"`
	}
	if err := c.ShouldBindJSON(&request); err != nil || len(request.TIDs) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "TIDs are required"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tickets booked successfully"})
}

// GetTicketStatus returns status of given tickets
func GetTicketStatus(c *gin.Context) {
	var request struct {
		TIDs []uint32 `json:"tids"`
	}
	if err := c.ShouldBindJSON(&request); err != nil || len(request.TIDs) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "TIDs are required"})
		return
	}

	// Dummy status mapping
	statusMap := make(map[uint32]string)
	for _, tid := range request.TIDs {
		statusMap[tid] = "Available"
	}

	c.JSON(http.StatusOK, statusMap)
}
