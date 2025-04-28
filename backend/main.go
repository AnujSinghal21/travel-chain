package main

import (
	"backend/middleware"
	"github.com/gin-gonic/gin"
	"backend/models"
	"backend/routes"
)

func main() {
	if err := middleware.InitializeFabric(); err != nil {
		// log.Fatalf("Failed to initialize Fabric connection: %v", err)
	}
	r := gin.Default()

	models.InitDB()

	// Register routes
	routes.RegisterRoutes(r)

	r.Run(":8080")
}
