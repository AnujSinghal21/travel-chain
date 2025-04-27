package controllers

import (
	"net/http"
	"fmt"
	"time"
	"backend/models"
	"github.com/gin-gonic/gin"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

func invokeChaincode(function string, args []string) (string, error) {
	sdk, err := fabsdk.New(config.FromFile("connection-org1.yaml"))
	if err != nil {
		return "", err
	}
	defer sdk.Close()

	clientContext := sdk.ChannelContext("mychannel", fabsdk.WithUser("Admin"), fabsdk.WithOrg("Org1"))
	client, err := channel.New(clientContext)
	if err != nil {
		return "", err
	}

	request := channel.Request{
		ChaincodeID: "mycc",
		Fcn:         function,
		Args:        args,
	}
	response, err := client.Execute(request)
	if err != nil {
		return "", err
	}
	return string(response.Payload), nil
}

func queryChaincode(function string, args []string) (string, error) {
	sdk, err := fabsdk.New(config.FromFile("connection-org1.yaml"))
	if err != nil {
		return "", err
	}
	defer sdk.Close()

	clientContext := sdk.ChannelContext("mychannel", fabsdk.WithUser("Admin"), fabsdk.WithOrg("Org1"))
	client, err := channel.New(clientContext)
	if err != nil {
		return "", err
	}

	request := channel.Request{
		ChaincodeID: "mycc",
		Fcn:         function,
		Args:        args,
	}
	response, err := client.Query(request)
	if err != nil {
		return "", err
	}
	return string(response.Payload), nil
}

// GetTicketByTID fetches a ticket by TID
func GetTicketByTID(c *gin.Context) {
	tid := c.Query("tid")
	if tid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "tid is required"})
		return
	}

	result, err := queryChaincode("getTicketByID", []string{tid})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get ticket", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// GetAllTickets returns all tickets
func GetAllTickets(c *gin.Context) {
	result, err := queryChaincode("getAllTickets", []string{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get tickets", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// CreateTickets creates tickets
func CreateTickets(c *gin.Context) {
	var tickets []models.Ticket
	if err := c.ShouldBindJSON(&tickets); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ticket list"})
		return
	}

	for _, ticket := range tickets {
		args := []string{
			fmt.Sprint(ticket.TID),
			ticket.ServiceID,
			fmt.Sprint(ticket.SeatNo),
			ticket.ServiceName,
			ticket.ServiceProviderID,
			fmt.Sprintf("%.2f", ticket.Price),
			ticket.StartTime.Format(time.RFC3339),
			ticket.Duration.String(),
			ticket.Source,
			ticket.Destination,
			ticket.TransportType,
		}
		_, err := invokeChaincode("createTicket", args)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create ticket", "details": err.Error()})
			return
		}
	}

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

	for _, tid := range request.TIDs {
		args := []string{fmt.Sprint(tid), "provider1@example.com"} // Replace with actual provider ID
		_, err := invokeChaincode("deleteTicket", args)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete ticket", "details": err.Error()})
			return
		}
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

	for _, tid := range request.TIDs {
		args := []string{"user1@example.com", fmt.Sprint(tid)} // Replace with actual user ID
		_, err := invokeChaincode("cancelTicket", args)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to cancel ticket", "details": err.Error()})
			return
		}
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

	for _, tid := range request.TIDs {
		args := []string{"user1@example.com", fmt.Sprint(tid)} // Replace with actual user ID
		_, err := invokeChaincode("bookTicket", args)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to book ticket", "details": err.Error()})
			return
		}
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

	statusMap := make(map[uint32]string)
	for _, tid := range request.TIDs {
		result, err := queryChaincode("getStatus", []string{fmt.Sprint(tid)})
		if err != nil {
			statusMap[tid] = "Error"
		} else {
			statusMap[tid] = result
		}
	}

	c.JSON(http.StatusOK, statusMap)
}