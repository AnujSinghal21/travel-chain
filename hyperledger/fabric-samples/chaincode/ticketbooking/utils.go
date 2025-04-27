package main

import (
	"strconv"
	"time"
	"github.com/hyperledger/fabric-chaincode-go/shim"
)

func parseUint32(s string) (uint32, error) {
	val, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint32(val), nil
}

func parseFloat(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

func getCurrentTime(stub shim.ChaincodeStubInterface) (time.Time, error) {
	timestamp, err := stub.GetTxTimestamp()
	if err != nil {
		return time.Time{}, err
	}
	return time.Unix(int64(timestamp.GetSeconds()), int64(timestamp.GetNanos())), nil
}