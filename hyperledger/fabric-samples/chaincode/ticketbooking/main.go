package main

import (
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/peer"
)

type Chaincode struct{}

func (cc *Chaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

func (cc *Chaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	function, args := stub.GetFunctionAndParameters()
	switch function {
	case "createUser":
		return createUser(stub, args)
	case "updateUser":
		return updateUser(stub, args)
	case "deleteUser":
		return deleteUser(stub, args)
	case "queryBalance":
		return queryBalance(stub, args)
	case "createTicket":
		return createTicket(stub, args)
	case "getTicketByID":
		return getTicketByID(stub, args)
	case "getAllTickets":
		return getAllTickets(stub, args)
	case "deleteTicket":
		return deleteTicket(stub, args)
	case "bookTicket":
		return bookTicket(stub, args)
	case "cancelTicket":
		return cancelTicket(stub, args)
	case "getStatus":
		return getStatus(stub, args)
	case "updateTicketPrice":
		return updateTicketPrice(stub, args)
	case "cancelListing":
		return cancelListing(stub, args)
	default:
		return shim.Error("Invalid function name")
	}
}

func main() {
	err := shim.Start(new(Chaincode))
	if err != nil {
		fmt.Printf("Error starting chaincode: %s", err)
	}
}