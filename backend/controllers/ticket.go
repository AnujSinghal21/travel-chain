package controllers

import (
	"backend/models"
	"net/http"
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetTicketByTID(c *gin.Context) {
	tid := c.Query("tid")
	if tid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ticket ID required"})
		return
	}

	result, err := queryChaincode("getTicketByID", []string{tid})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to retrieve ticket",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"ticket": result})
}


func GetAllTickets(c *gin.Context) {
	result, err := queryChaincode("getAllTickets", []string{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to retrieve tickets",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tickets": result})
}

// CreateTicket by provider
func CreateTickets(c *gin.Context) {
	var tickets []models.Ticket
	if err := c.ShouldBindJSON(&tickets); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ticket data"})
		return
	}

	email, _ := c.Get("email")
	for _, ticket := range tickets {
		args := []string{
			fmt.Sprint(ticket.TID),
			ticket.ServiceID,
			fmt.Sprint(ticket.SeatNo),
			ticket.ServiceName,
			email.(string), 
			fmt.Sprintf("%.2f", ticket.Price),
			ticket.StartTime,
			fmt.Sprint(ticket.Duration),
			ticket.Source,
			ticket.Destination,
			ticket.TransportType,
		}

		if _, err := invokeChaincode("createTicket", args); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Failed to create ticket",
				"details": err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tickets created successfully"})
}

func BookTickets(c *gin.Context) {
	var request struct {
		TIDs []uint32 `json:"tids"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	email, _ := c.Get("email")
	for _, tid := range request.TIDs {
		args := []string{
			email.(string),
			fmt.Sprint(tid),
		}

		if _, err := invokeChaincode("bookTicket", args); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Failed to book ticket",
				"details": err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tickets booked successfully"})
}

func CancelTickets(c *gin.Context) {
	var request struct {
		TIDs []uint32 `json:"tids"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	email, _ := c.Get("email")
	for _, tid := range request.TIDs {
		args := []string{
			email.(string),
			fmt.Sprint(tid),
		}

		if _, err := invokeChaincode("cancelTicket", args); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Failed to cancel ticket",
				"details": err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tickets cancelled successfully"})
}

func GetTicketStatus(c *gin.Context) {
	var request struct {
		TIDs []uint32 `json:"tids"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	statusMap := make(map[uint32]string)
	for _, tid := range request.TIDs {
		result, err := queryChaincode("getStatus", []string{fmt.Sprint(tid)})
		if err != nil {
			statusMap[tid] = "error"
		} else {
			statusMap[tid] = result
		}
	}

	c.JSON(http.StatusOK, gin.H{"status": statusMap})
}
