package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
	"github.com/anujSinghal21/travel-chain/api-interface/services"
)

func GetUserDetails(c *gin.Context) {
	certificate := c.Query("certificate")
	if certificate == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Certificate is required"})
		return
	}

	userDetails, err := services.GetUserDetails(certificate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user details"})
		return
	}

	c.JSON(http.StatusOK, userDetails)
}