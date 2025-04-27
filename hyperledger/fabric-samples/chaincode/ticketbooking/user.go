package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/peer"
)

func createUser(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 5 {
		return shim.Error("Expecting 5 args: email, name, role, password, balance")
	}
	email := args[0]
	name := args[1]
	role := args[2]
	balance, err := parseFloat(args[4])
	if err != nil {
		return shim.Error("Invalid balance")
	}
	userKey := "USER_" + email
	userAsBytes, _ := stub.GetState(userKey)
	if userAsBytes != nil {
		return shim.Error("User already exists")
	}
	user := User{DocType: "user", Email: email, Name: name, Role: role, Balance: balance}
	userJSON, _ := json.Marshal(user)
	err = stub.PutState(userKey, userJSON)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

func queryBalance(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("Expecting 1 arg: email")
	}
	email := args[0]
	userKey := "USER_" + email
	userAsBytes, err := stub.GetState(userKey)
	if err != nil || userAsBytes == nil {
		return shim.Error("User not found")
	}
	var user User
	json.Unmarshal(userAsBytes, &user)
	return shim.Success([]byte(fmt.Sprintf("%f", user.Balance)))
}

func updateUser(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 4 {
		return shim.Error("Expecting 4 args: email, requesterEmail, role, balance")
	}
	email := args[0]
	requesterEmail := args[1]
	role := args[2]
	balance, err := parseFloat(args[3])
	if err != nil {
		return shim.Error("Invalid balance")
	}
	if email != requesterEmail {
		return shim.Error("Only the user can update their own details")
	}
	userKey := "USER_" + email
	userAsBytes, err := stub.GetState(userKey)
	if err != nil || userAsBytes == nil {
		return shim.Error("User not found")
	}
	var user User
	json.Unmarshal(userAsBytes, &user)
	user.Role = role
	user.Balance = balance
	userJSON, _ := json.Marshal(user)
	err = stub.PutState(userKey, userJSON)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

func deleteUser(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 2 {
		return shim.Error("Expecting 2 args: email, requesterEmail")
	}
	email := args[0]
	requesterEmail := args[1]
	if email != requesterEmail {
		return shim.Error("Only the user can delete their own account")
	}
	userKey := "USER_" + email
	userAsBytes, err := stub.GetState(userKey)
	if err != nil || userAsBytes == nil {
		return shim.Error("User not found")
	}

	queryString := fmt.Sprintf(`{"selector":{"docType":"ticket","passenger":"%s","status":"booked"}}`, email)
	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()
	if resultsIterator.HasNext() {
		return shim.Error("Cannot delete user with booked tickets")
	}
	err = stub.DelState(userKey)
	if err != nil {
		return shim.Error("Failed to delete user")
	}
	return shim.Success(nil)
}