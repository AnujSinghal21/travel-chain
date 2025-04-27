package main

import (
    "encoding/json"
    "fmt"
    "strconv"
    "time"

    "github.com/hyperledger/fabric-chaincode-go/shim"
    "github.com/hyperledger/fabric-protos-go/peer"
)

// createTicket creates a new ticket with seat uniqueness check
func createTicket(stub shim.ChaincodeStubInterface, args []string) peer.Response {
    if len(args) != 11 {
        return shim.Error("Expecting 11 args: tid, serviceID, seatNo, serviceName, serviceProviderID, price, startTime, duration, source, destination, transportType")
    }
    tid, err := parseUint32(args[0])
    if err != nil {
        return shim.Error("Invalid TID")
    }
    serviceID := args[1]
    seatNo, err := parseUint32(args[2])
    if err != nil {
        return shim.Error("Invalid seat number")
    }
    serviceName := args[3]
    serviceProviderID := args[4]
    price, err := parseFloat(args[5])
    if err != nil {
        return shim.Error("Invalid price")
    }
    startTime, err := time.Parse(time.RFC3339, args[6])
    if err != nil {
        return shim.Error("Invalid start time format (use RFC3339)")
    }
    duration, err := time.ParseDuration(args[7])
    if err != nil {
        return shim.Error("Invalid duration")
    }
    source := args[8]
    destination := args[9]
    transportType := args[10]

    // Check seat uniqueness
    seatKey, err := stub.CreateCompositeKey("seat", []string{serviceID, fmt.Sprint(seatNo)})
    if err != nil {
        return shim.Error("Failed to create composite key")
    }
    seatExists, err := stub.GetState(seatKey)
    if err != nil {
        return shim.Error("Failed to check seat: " + err.Error())
    }
    if seatExists != nil {
        return shim.Error("Seat already exists in this service")
    }

    // Verify provider
    providerKey := "USER_" + serviceProviderID
    providerAsBytes, err := stub.GetState(providerKey)
    if err != nil || providerAsBytes == nil {
        return shim.Error("Provider does not exist")
    }
    var provider User
    if err := json.Unmarshal(providerAsBytes, &provider); err != nil {
        return shim.Error("Failed to unmarshal provider: " + err.Error())
    }
    if provider.Role != "provider" {
        return shim.Error("Only providers can create tickets")
    }

    ticketKey := fmt.Sprintf("TICKET_%d", tid)
    ticketAsBytes, err := stub.GetState(ticketKey)
    if err != nil {
        return shim.Error("Failed to check ticket: " + err.Error())
    }
    if ticketAsBytes != nil {
        return shim.Error("Ticket already exists")
    }

    ticket := Ticket{
        DocType:           "ticket",
        TID:               tid,
        ServiceID:         serviceID,
        SeatNo:            seatNo,
        ServiceName:       serviceName,
        ServiceProviderID: serviceProviderID,
        Status:            "available",
        PassengerID:       "",
        Price:             price,
        StartTime:         startTime,
        Duration:          duration,
        Source:            source,
        Destination:       destination,
        TransportType:     transportType,
    }
    ticketJSON, err := json.Marshal(ticket)
    if err != nil {
        return shim.Error("Failed to marshal ticket: " + err.Error())
    }
    if err = stub.PutState(ticketKey, ticketJSON); err != nil {
        return shim.Error("Failed to save ticket: " + err.Error())
    }
    // Mark seat as taken
    if err = stub.PutState(seatKey, []byte{0x00}); err != nil {
        return shim.Error("Failed to mark seat: " + err.Error())
    }
    return shim.Success(nil)
}

// getTicketByID retrieves a ticket by its TID
func getTicketByID(stub shim.ChaincodeStubInterface, args []string) peer.Response {
    if len(args) != 1 {
        return shim.Error("Expecting 1 arg: tid")
    }
    tid, err := parseUint32(args[0])
    if err != nil {
        return shim.Error("Invalid TID")
    }
    ticketKey := fmt.Sprintf("TICKET_%d", tid)
    ticketAsBytes, err := stub.GetState(ticketKey)
    if err != nil || ticketAsBytes == nil {
        return shim.Error("Ticket not found")
    }
    return shim.Success(ticketAsBytes)
}

// getAllTickets retrieves all tickets for frontend filtering
func getAllTickets(stub shim.ChaincodeStubInterface, args []string) peer.Response {
    queryString := `{"selector":{"docType":"ticket"}}`
    resultsIterator, err := stub.GetQueryResult(queryString)
    if err != nil {
        return shim.Error("Failed to query tickets: " + err.Error())
    }
    defer resultsIterator.Close()
    var tickets []Ticket
    for resultsIterator.HasNext() {
        queryResponse, err := resultsIterator.Next()
        if err != nil {
            return shim.Error("Iterator error: " + err.Error())
        }
        var ticket Ticket
        if err := json.Unmarshal(queryResponse.Value, &ticket); err != nil {
            return shim.Error("Failed to unmarshal ticket: " + err.Error())
        }
        tickets = append(tickets, ticket)
    }
    ticketsJSON, err := json.Marshal(tickets)
    if err != nil {
        return shim.Error("Failed to marshal tickets: " + err.Error())
    }
    return shim.Success(ticketsJSON)
}

// deleteTicket deletes a ticket if it is available
func deleteTicket(stub shim.ChaincodeStubInterface, args []string) peer.Response {
    if len(args) != 2 {
        return shim.Error("Expecting 2 args: tid, serviceProviderID")
    }
    tid, err := parseUint32(args[0])
    if err != nil {
        return shim.Error("Invalid TID")
    }
    serviceProviderID := args[1]
    ticketKey := fmt.Sprintf("TICKET_%d", tid)
    ticketAsBytes, err := stub.GetState(ticketKey)
    if err != nil || ticketAsBytes == nil {
        return shim.Error("Ticket not found")
    }
    var ticket Ticket
    if err := json.Unmarshal(ticketAsBytes, &ticket); err != nil {
        return shim.Error("Failed to unmarshal ticket: " + err.Error())
    }
    if ticket.ServiceProviderID != serviceProviderID {
        return shim.Error("Only ticket provider can delete")
    }
    if ticket.Status != "available" {
        return shim.Error("Only available tickets can be deleted")
    }
    // Remove seat marker
    seatKey, err := stub.CreateCompositeKey("seat", []string{ticket.ServiceID, fmt.Sprint(ticket.SeatNo)})
    if err != nil {
        return shim.Error("Failed to create composite key")
    }
    if err = stub.DelState(seatKey); err != nil {
        return shim.Error("Failed to remove seat marker: " + err.Error())
    }
    if err = stub.DelState(ticketKey); err != nil {
        return shim.Error("Failed to delete ticket: " + err.Error())
    }
    return shim.Success(nil)
}

// bookTicket books a ticket for a user
func bookTicket(stub shim.ChaincodeStubInterface, args []string) peer.Response {
    if len(args) != 2 {
        return shim.Error("Expecting 2 args: userID, tid")
    }
    userID := args[0]
    tid, err := parseUint32(args[1])
    if err != nil {
        return shim.Error("Invalid TID")
    }
    ticketKey := fmt.Sprintf("TICKET_%d", tid)
    userKey := "USER_" + userID
    confirmationKey := fmt.Sprintf("CONFIRMATION_%d", tid)

    ticketAsBytes, err := stub.GetState(ticketKey)
    if err != nil || ticketAsBytes == nil {
        return shim.Error("Ticket not found")
    }
    var ticket Ticket
    if err := json.Unmarshal(ticketAsBytes, &ticket); err != nil {
        return shim.Error("Failed to unmarshal ticket: " + err.Error())
    }
    if ticket.Status != "available" {
        return shim.Error("Ticket not available")
    }

    userAsBytes, err := stub.GetState(userKey)
    if err != nil || userAsBytes == nil {
        return shim.Error("User not found")
    }
    var user User
    if err := json.Unmarshal(userAsBytes, &user); err != nil {
        return shim.Error("Failed to unmarshal user: " + err.Error())
    }
    if user.Role != "user" {
        return shim.Error("Only users can book tickets")
    }
    if user.Balance < ticket.Price {
        return shim.Error("Insufficient balance")
    }

    ticket.Status = "booked"
    ticket.PassengerID = userID
    user.Balance -= ticket.Price

    ticketJSON, err := json.Marshal(ticket)
    if err != nil {
        return shim.Error("Failed to marshal ticket: " + err.Error())
    }
    if err = stub.PutState(ticketKey, ticketJSON); err != nil {
        return shim.Error("Failed to save ticket: " + err.Error())
    }
    userJSON, err := json.Marshal(user)
    if err != nil {
        return shim.Error("Failed to marshal user: " + err.Error())
    }
    if err = stub.PutState(userKey, userJSON); err != nil {
        return shim.Error("Failed to save user: " + err.Error())
    }

    timestamp, err := stub.GetTxTimestamp()
    if err != nil {
        return shim.Error("Failed to get timestamp: " + err.Error())
    }
    blockHeight := int(timestamp.GetSeconds())
    if err = stub.PutState(confirmationKey, []byte(fmt.Sprintf("%d", blockHeight))); err != nil {
        return shim.Error("Failed to save confirmation: " + err.Error())
    }

    return shim.Success([]byte("Ticket booked successfully"))
}

// cancelTicket cancels a ticket with pro-rata refund
func cancelTicket(stub shim.ChaincodeStubInterface, args []string) peer.Response {
    if len(args) != 2 {
        return shim.Error("Expecting 2 args: userID, tid")
    }
    userID := args[0]
    tid, err := parseUint32(args[1])
    if err != nil {
        return shim.Error("Invalid TID")
    }
    ticketKey := fmt.Sprintf("TICKET_%d", tid)
    userKey := "USER_" + userID
    confirmationKey := fmt.Sprintf("CONFIRMATION_%d", tid)

    ticketAsBytes, err := stub.GetState(ticketKey)
    if err != nil || ticketAsBytes == nil {
        return shim.Error("Ticket not found")
    }
    var ticket Ticket
    if err := json.Unmarshal(ticketAsBytes, &ticket); err != nil {
        return shim.Error("Failed to unmarshal ticket: " + err.Error())
    }
    if ticket.PassengerID != userID || ticket.Status != "booked" {
        return shim.Error("Cannot cancel this ticket")
    }

    userAsBytes, err := stub.GetState(userKey)
    if err != nil || userAsBytes == nil {
        return shim.Error("User not found")
    }
    var user User
    if err := json.Unmarshal(userAsBytes, &user); err != nil {
        return shim.Error("Failed to unmarshal user: " + err.Error())
    }

    currentTime, err := getCurrentTime(stub)
    if err != nil {
        return shim.Error("Failed to get current time: " + err.Error())
    }
    daysUntilTravel := int(ticket.StartTime.Sub(currentTime).Hours() / 24)
    var refundPercentage float64
    if daysUntilTravel > 2 {
        refundPercentage = 1.0
    } else if daysUntilTravel >= 0 {
        refundPercentage = 0.5
    } else {
        return shim.Error("Cannot cancel past travel")
    }

    refundAmount := ticket.Price * refundPercentage
    ticket.Status = "available"
    ticket.PassengerID = ""

    // Remove seat marker
    seatKey, err := stub.CreateCompositeKey("seat", []string{ticket.ServiceID, fmt.Sprint(ticket.SeatNo)})
    if err != nil {
        return shim.Error("Failed to create composite key")
    }
    if err = stub.DelState(seatKey); err != nil {
        return shim.Error("Failed to remove seat marker: " + err.Error())
    }

    ticketJSON, err := json.Marshal(ticket)
    if err != nil {
        return shim.Error("Failed to marshal ticket: " + err.Error())
    }
    if err = stub.PutState(ticketKey, ticketJSON); err != nil {
        return shim.Error("Failed to save ticket: " + err.Error())
    }
    user.Balance += refundAmount
    userJSON, err := json.Marshal(user)
    if err != nil {
        return shim.Error("Failed to marshal user: " + err.Error())
    }
    if err = stub.PutState(userKey, userJSON); err != nil {
        return shim.Error("Failed to save user: " + err.Error())
    }

    // Remove confirmation
    if _, err := stub.GetState(confirmationKey); err == nil {
        if err = stub.DelState(confirmationKey); err != nil {
            return shim.Error("Failed to remove confirmation: " + err.Error())
        }
    }

    return shim.Success([]byte(fmt.Sprintf("Ticket cancelled with %.2f refund", refundAmount)))
}

// getStatus checks the status of a ticket, including 2-block confirmation
func getStatus(stub shim.ChaincodeStubInterface, args []string) peer.Response {
    if len(args) != 1 {
        return shim.Error("Expecting 1 arg: tid")
    }
    tid, err := parseUint32(args[0])
    if err != nil {
        return shim.Error("Invalid TID")
    }
    ticketKey := fmt.Sprintf("TICKET_%d", tid)
    confirmationKey := fmt.Sprintf("CONFIRMATION_%d", tid)

    ticketAsBytes, err := stub.GetState(ticketKey)
    if err != nil || ticketAsBytes == nil {
        return shim.Error("Ticket not found")
    }
    var ticket Ticket
    if err := json.Unmarshal(ticketAsBytes, &ticket); err != nil {
        return shim.Error("Failed to unmarshal ticket: " + err.Error())
    }

    if ticket.Status != "booked" {
        return shim.Success([]byte(fmt.Sprintf("Status: %s", ticket.Status)))
    }

    confirmationAsBytes, err := stub.GetState(confirmationKey)
    if err != nil || confirmationAsBytes == nil {
        return shim.Error("Confirmation data not found")
    }
    blockHeight, err := strconv.Atoi(string(confirmationAsBytes))
    if err != nil {
        return shim.Error("Invalid block height")
    }

    timestamp, err := stub.GetTxTimestamp()
    if err != nil {
        return shim.Error("Failed to get timestamp: " + err.Error())
    }
    currentHeight := int(timestamp.GetSeconds())
    if currentHeight >= blockHeight+2 {
        return shim.Success([]byte("Ticket confirmed"))
    }
    return shim.Success([]byte("Status: booked (pending confirmation)"))
}

// updateTicketPrice updates the price with a 10% surge when >50% seats are already booked
func updateTicketPrice(stub shim.ChaincodeStubInterface, args []string) peer.Response {
    if len(args) != 3 {
        return shim.Error("Expecting 3 args: tid, serviceProviderID, newPrice")
    }
    tid, err := parseUint32(args[0])
    if err != nil {
        return shim.Error("Invalid TID")
    }
    providerID := args[1]
    basePrice, err := parseFloat(args[2])
    if err != nil {
        return shim.Error("Invalid price")
    }

    // Fetch ticket
    ticketKey := fmt.Sprintf("TICKET_%d", tid)
    ticketAsBytes, err := stub.GetState(ticketKey)
    if err != nil || ticketAsBytes == nil {
        return shim.Error("Ticket not found")
    }
    var ticket Ticket
    if err := json.Unmarshal(ticketAsBytes, &ticket); err != nil {
        return shim.Error("Failed to unmarshal ticket: " + err.Error())
    }

    // Authorization & status checks
    if ticket.ServiceProviderID != providerID {
        return shim.Error("Only ticket provider can update price")
    }
    if ticket.Status != "available" {
        return shim.Error("Only available tickets can update price")
    }

    // Count total & booked tickets for this service
    query := fmt.Sprintf(`{"selector":{"docType":"ticket","service_id":"%s"}}`, ticket.ServiceID)
    resultsIterator, err := stub.GetQueryResult(query)
    if err != nil {
        return shim.Error("Failed to query tickets: " + err.Error())
    }
    defer resultsIterator.Close()

    var total, booked int
    for resultsIterator.HasNext() {
        queryResponse, err := resultsIterator.Next()
        if err != nil {
            return shim.Error("Iterator error: " + err.Error())
        }
        var t Ticket
        if err := json.Unmarshal(queryResponse.Value, &t); err != nil {
            return shim.Error("Failed to unmarshal ticket: " + err.Error())
        }
        total++
        if t.Status == "booked" {
            booked++
        }
    }

    // Apply 10% increase if more than half seats booked
    finalPrice := basePrice
    if total > 0 && float64(booked)/float64(total) > 0.5 {
        finalPrice = basePrice * 1.1
    }

    // Persist updated price
    ticket.Price = finalPrice
    updated, err := json.Marshal(ticket)
    if err != nil {
        return shim.Error("Failed to marshal ticket: " + err.Error())
    }
    if err := stub.PutState(ticketKey, updated); err != nil {
        return shim.Error("Failed to save ticket: " + err.Error())
    }

    return shim.Success([]byte(fmt.Sprintf("Price updated to %.2f", finalPrice)))
}

// cancelListing allows a provider to cancel an entire journey (serviceID)
func cancelListing(stub shim.ChaincodeStubInterface, args []string) peer.Response {
    if len(args) != 2 {
        return shim.Error("Expecting 2 args: serviceProviderID, serviceID")
    }
    serviceProviderID := args[0]
    serviceID := args[1]

    // Verify provider exists and has correct role
    providerKey := "USER_" + serviceProviderID
    providerAsBytes, err := stub.GetState(providerKey)
    if err != nil || providerAsBytes == nil {
        return shim.Error("Provider not found")
    }
    var provider User
    if err := json.Unmarshal(providerAsBytes, &provider); err != nil {
        return shim.Error("Failed to unmarshal provider: " + err.Error())
    }
    if provider.Role != "provider" {
        return shim.Error("Only providers can cancel listings")
    }

    // Query all tickets for this service
    query := fmt.Sprintf(`{"selector":{"docType":"ticket","service_id":"%s"}}`, serviceID)
    resultsIterator, err := stub.GetQueryResult(query)
    if err != nil {
        return shim.Error("Failed to query tickets: " + err.Error())
    }
    defer resultsIterator.Close()

    // Process each ticket: refund if booked, delete state
    for resultsIterator.HasNext() {
        queryResponse, err := resultsIterator.Next()
        if err != nil {
            return shim.Error("Iterator error: " + err.Error())
        }
        var ticket Ticket
        if err := json.Unmarshal(queryResponse.Value, &ticket); err != nil {
            return shim.Error("Failed to unmarshal ticket: " + err.Error())
        }

        // Refund if booked
        if ticket.Status == "booked" {
            userKey := "USER_" + ticket.PassengerID
            userAsBytes, err := stub.GetState(userKey)
            if err != nil || userAsBytes == nil {
                return shim.Error(fmt.Sprintf("User %s not found for ticket %d", ticket.PassengerID, ticket.TID))
            }
            var user User
            if err := json.Unmarshal(userAsBytes, &user); err != nil {
                return shim.Error("Failed to unmarshal user: " + err.Error())
            }
            user.Balance += ticket.Price
            userJSON, err := json.Marshal(user)
            if err != nil {
                return shim.Error("Failed to marshal user: " + err.Error())
            }
            if err := stub.PutState(userKey, userJSON); err != nil {
                return shim.Error("Failed to save user: " + err.Error())
            }
        }

        // Remove confirmation entry
        confirmationKey := fmt.Sprintf("CONFIRMATION_%d", ticket.TID)
        if _, err := stub.GetState(confirmationKey); err == nil {
            if err := stub.DelState(confirmationKey); err != nil {
                return shim.Error("Failed to remove confirmation: " + err.Error())
            }
        }

        // Remove seat marker
        seatKey, err := stub.CreateCompositeKey("seat", []string{ticket.ServiceID, fmt.Sprint(ticket.SeatNo)})
        if err != nil {
            return shim.Error("Failed to create composite key")
        }
        if err := stub.DelState(seatKey); err != nil {
            return shim.Error("Failed to remove seat marker: " + err.Error())
        }

        // Delete the ticket
        ticketKey := fmt.Sprintf("TICKET_%d", ticket.TID)
        if err := stub.DelState(ticketKey); err != nil {
            return shim.Error("Failed to delete ticket: " + err.Error())
        }
    }

    return shim.Success([]byte("Journey cancelled; all tickets removed and refunds processed"))
}