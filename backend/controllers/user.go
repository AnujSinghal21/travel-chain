package controllers

import (
	"backend/models"
	"net/http"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

const (
	chaincodeID = "mycc" 
	channelID   = "mychannel"
)

var (
	sdk *fabsdk.FabricSDK
)

func initFabric() error {
	if sdk == nil {
		configPath := "../hyperledger/fabric-samples/test-network/organizations/peerOrganizations/org1.example.com/connection-org1.yaml"
		var err error
		sdk, err = fabsdk.New(config.FromFile(configPath))
		if err != nil {
			return fmt.Errorf("failed to create Fabric SDK: %v", err)
		}
	}
	return nil
}

func invokeChaincode(function string, args []string) (string, error) {
	if err := initFabric(); err != nil {
		return "", err
	}

	clientChannelContext := sdk.ChannelContext(channelID, fabsdk.WithUser("Admin"), fabsdk.WithOrg("Org1"))
	client, err := channel.New(clientChannelContext)
	if err != nil {
		return "", fmt.Errorf("failed to create channel client: %v", err)
	}

	response, err := client.Execute(channel.Request{
		ChaincodeID: chaincodeID,
		Fcn:         function,
		Args:        convertArgs(args),
	})
	if err != nil {
		return "", fmt.Errorf("chaincode invocation failed: %v", err)
	}

	return string(response.Payload), nil
}

func queryChaincode(function string, args []string) (string, error) {
	if err := initFabric(); err != nil {
		return "", err
	}

	clientChannelContext := sdk.ChannelContext(channelID, fabsdk.WithUser("Admin"), fabsdk.WithOrg("Org1"))
	client, err := channel.New(clientChannelContext)
	if err != nil {
		return "", fmt.Errorf("failed to create channel client: %v", err)
	}

	response, err := client.Query(channel.Request{
		ChaincodeID: chaincodeID,
		Fcn:         function,
		Args:        convertArgs(args),
	})
	if err != nil {
		return "", fmt.Errorf("chaincode query failed: %v", err)
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

func RegisterHandler(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	requiredFields := []string{user.Email, user.Password, user.Name, user.Phone, user.Role}
	for _, field := range requiredFields {
		if field == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "All fields are required"})
			return
		}
	}


	hashedPassword := user.Password 

	args := []string{
		user.Email,
		user.Name,
		user.Role,
		hashedPassword,
		"1000.00", // balance
	}

	if _, err := invokeChaincode("createUser", args); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Blockchain registration failed",
			"details": err.Error(),
		})
		c.Error(err)
		return
	}

	user.Password = hashedPassword
	if err := models.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database operation failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}


func LoginHandler(c *gin.Context) {
	var loginReq models.LoginRequest
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	var user models.User
	if err := models.DB.Where("email = ?", loginReq.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if user.Password != loginReq.Password { // Use bcrypt compare in production
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := models.GenerateToken(user.Email, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user": gin.H{
			"email": user.Email,
			"name":  user.Name,
			"role":  user.Role,
		},
	})
}
n
func GetUserDetailsHandler(c *gin.Context) {
	email, exists := c.Get("email")
	if !exists || email == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	balance, err := queryChaincode("queryBalance", []string{email.(string)})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to retrieve balance",
			"details": err.Error(),
		})
		return
	}

	var user models.User
	if err := models.DB.Where("email = ?", email).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"name":    user.Name,
		"email":   user.Email,
		"phone":   user.Phone,
		"role":    user.Role,
		"balance": balance,
	})
}

func LogoutHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}
