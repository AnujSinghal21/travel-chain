package controllers

import (
	"net/http"

	"backend/models"

	"github.com/gin-gonic/gin"
)

// RegisterHandler handles user registration
func RegisterHandler(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		c.Error(err) // Print debug message to show the error in binding
		return
	}

	if user.Email == "" || user.Password == "" || user.Name == "" || user.Phone == "" || user.Role == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields"})
		return
	}

	// Save user to DB
	if err := models.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists or database error"})
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
		c.Error(err) // Print debug message to show the error in binding
		return
	}

	// Find user by email
	if err := models.DB.Where("email = ?", loginReq.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	if user.Password != loginReq.Password {
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

	var user models.User
	if err := models.DB.Where("email = ?", email).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"name":  user.Name,
		"email": user.Email,
		"phone": user.Phone,
		"role":  user.Role,
		"age":   user.Age,
	})
}
