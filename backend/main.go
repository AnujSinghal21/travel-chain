package main

import (
	"github.com/gin-gonic/gin"
	"backend/models"
	"backend/routes"
)

func main() {
	r := gin.Default()

	// Initialize database
	models.InitDB()

	// Register routes
	routes.RegisterRoutes(r)

	r.Run(":8080")
}
