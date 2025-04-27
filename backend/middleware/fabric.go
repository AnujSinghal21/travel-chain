package middleware

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

var (
	sdkInstance *fabsdk.FabricSDK
	channelName = "mychannel"
	ccName      = "ticketcc"
)

func InitializeFabric() error {
	configPath := filepath.Join("..", "Hyperledger", "fabric-samples", "test-network", "organizations", "peerOrganizations", "org1.example.com", "connection-org1.yaml")
	
	sdk, err := fabsdk.New(config.FromFile(configPath))
	if err != nil {
		return fmt.Errorf("failed to create SDK: %v", err)
	}

	sdkInstance = sdk
	return nil
}

func getClient() (*channel.Client, error) {
	context := sdkInstance.ChannelContext(
		channelName,
		fabsdk.WithUser("Admin"),
		fabsdk.WithOrg("Org1"),
	)

	return channel.New(context)
}

func InvokeChaincode(fn string, args []string) (string, error) {
	client, err := getClient()
	if err != nil {
		return "", err
	}

	response, err := client.Execute(channel.Request{
		ChaincodeID: ccName,
		Fcn:         fn,
		Args:        convertArgs(args),
	}, channel.WithTimeout(fab.Timeout, 30*time.Second))

	if err != nil {
		return "", fmt.Errorf("invoke error: %v", err)
	}

	return string(response.Payload), nil
}

func QueryChaincode(fn string, args []string) (string, error) {
	client, err := getClient()
	if err != nil {
		return "", err
	}

	response, err := client.Query(channel.Request{
		ChaincodeID: ccName,
		Fcn:         fn,
		Args:        convertArgs(args),
	})

	if err != nil {
		return "", fmt.Errorf("query error: %v", err)
	}

	return string(response.Payload), nil
}

func convertArgs(args []string) [][]byte {
	byteArgs := make([][]byte, len(args))
	for i, arg := range args {
		byteArgs[i] = []byte(arg)
	}
	return byteArgs
}
