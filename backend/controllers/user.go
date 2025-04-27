package controllers

import (
	"backend/models"
	"net/http"

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

// RegisterHandler handles user registration
func RegisterHandler(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if user.Email == "" || user.Password == "" || user.Name == "" || user.Phone == "" || user.Role == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields"})
		return
	}

	// Hash password and store in local DB (assuming a hashing function exists)
	hashedPassword := user.Password // Replace with actual hashing logic

	// Call chaincode to create user
	args := []string{user.Email, user.Name, user.Role, hashedPassword, "1000.00"} // Initial balance
	_, err := invokeChaincode("createUser", args)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user on blockchain", "details": err.Error()})
		return
	}

	// Optionally save to local DB for authentication
	user.Password = hashedPassword
	if err := models.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Database error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
}

// LoginHandler handles user login
func LoginHandler(c *gin.Context) {
	var loginReq models.LoginRequest
	var user models.User

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	// Find user by email in local DB
	if err := models.DB.Where("email = ?", loginReq.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	if user.Password != loginReq.Password { // Replace with proper password check
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Generate JWT token
	token, err := models.GenerateToken(user.Email, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token, "user": user})
}

// LogoutHandler handles user logout
func LogoutHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}

// GetUserDetailsHandler returns user profile from token
func GetUserDetailsHandler(c *gin.Context) {
	emailValue, exists := c.Get("email")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	email, ok := emailValue.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid token data"})
		return
	}

	// Query balance from chaincode
	balance, err := queryChaincode("queryBalance", []string{email})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query balance", "details": err.Error()})
		return
	}

	// Fetch user details from local DB
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
		"age":     user.Age,
		"balance": balance,
	})
}
