package middleware

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/logging"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/context"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

const (
	channelName = "mychannel"
	chaincodeID = "mycc"
	userName    = "Admin"
	orgName     = "Org1"
)

var (
	sdk      *fabsdk.FabricSDK
	chClient *channel.Client
)

func InitializeFabric() error {
	logging.SetLevel("", logging.DEBUG)
	configPath := getConfigPath()
	fmt.Printf("Using config path: %s\n", configPath)
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return fmt.Errorf("config file does not exist: %s", configPath)
	}
	configBackend := config.FromFile(configPath)
	fmt.Println("Config backend created")
	sdk, err := fabsdk.New(configBackend)
	if err != nil {
		return fmt.Errorf("failed to create SDK: %v", err)
	}
	fmt.Println("SDK created successfully")

	mspClient, err := msp.New(
		sdk.Context(
			fabsdk.WithOrg("Org1"),
		),
	)
	if err != nil {

		return fmt.Errorf("Failed to create MSP client: %v", err)
	}
	fmt.Println("mspClient created successfully")

	// Enroll admin if not enrolled
	_, err = mspClient.GetSigningIdentity("Admin")
	if err != nil {
		fmt.Println("Admin not enrolled yet. Enrolling now...")
		err = mspClient.Enroll("Admin", msp.WithSecret("adminpw"), msp.WithLabel("Admin"))
		if err != nil {
			log.Fatalf("Failed to enroll admin: %v", err)
		}
		fmt.Println("Admin enrolled successfully.")
	} else {

		fmt.Println("Admin already enrolled.")
	}
	id, err := mspClient.GetSigningIdentity("Admin")
	if err != nil {
		log.Printf("Failed to get signing identity: %v", err)
	} else {
		log.Printf("Successfully loaded identity: %s", id.Identifier().ID)
	}

	chContext := sdk.ChannelContext(
		channelName,
		fabsdk.WithUser(userName),
		fabsdk.WithOrg(orgName),
	)
	chClient, err = channel.New(chContext)
	if err != nil {
		return fmt.Errorf("failed to create channel client: %v", err)
	}
	fmt.Println("Channel client created successfully")
	return nil
}

func InvokeChaincode(fn string, args []string) (string, error) {
	request := channel.Request{
		ChaincodeID: chaincodeID,
		Fcn:         fn,
		Args:        convertArgs(args),
	}

	response, err := chClient.Execute(request, channel.WithTimeout(fab.PeerResponse, 30*time.Second))
	if err != nil {
		return "", fmt.Errorf("invoke failed: %v", err)
	}

	return string(response.Payload), nil
}

func QueryChaincode(fn string, args []string) (string, error) {
	request := channel.Request{
		ChaincodeID: chaincodeID,
		Fcn:         fn,
		Args:        convertArgs(args),
	}

	response, err := chClient.Query(request)
	if err != nil {
		return "", fmt.Errorf("query failed: %v", err)
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

// func getConfigPath() string {
// 	return filepath.Join(
// 		getProjectRoot(),
// 		"hyperledger",
// 		"fabric-samples",
// 		"test-network",
// 		"organizations",
// 		"peerOrganizations",
// 		"org1.example.com",
// 		"connection-org1.yaml",
// 	)
// }

func getConfigPath() string {
	return "/home/goutam/TicketBookingSystem/hyperledger/fabric-samples/test-network/organizations/peerOrganizations/org1.example.com/connection-org1.yaml"
	// return "/home/goutam/TicketBookingSystem/backend/config/connection-org1.yaml"
}

func getOrgPath() string {
	return filepath.Join(
		getProjectRoot(),
		"hyperledger",
		"fabric-samples",
		"test-network",
		"organizations",
		"peerOrganizations",
		"org1.example.com",
	)
} // bekar

func getProjectRoot() string {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = filepath.Join(os.Getenv("HOME"), "go")
	}
	return filepath.Join(gopath, "src", "TicketBookingSystem")
}

func GetChannelContext() (context.ChannelProvider, error) {
	if sdk == nil {
		return nil, fmt.Errorf("SDK not initialized")
	}
	return sdk.ChannelContext(
		channelName,
		fabsdk.WithUser(userName),
		fabsdk.WithOrg(orgName),
	), nil
}

// GetSigningIdentity returns the organization's signing identity
// func GetSigningIdentity() (msp.SigningIdentity, error) {
// 	if sdk == nil {
// 		return nil, fmt.Errorf("SDK not initialized")
// 	}

// 	ctxProvider := sdk.Context()
// 	ctx, err := ctxProvider()
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to get context: %v", err)
// 	}

// 	identityManager, ok := ctx.IdentityManager(orgName)
// 	if !ok {
// 		return nil, fmt.Errorf("failed to get identity manager for org: %s", orgName)
// 	}
// 	return identityManager.GetSigningIdentity(userName)
// }
