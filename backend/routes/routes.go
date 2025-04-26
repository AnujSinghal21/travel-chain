package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/anujSinghal21/travel-chain/api-interface/controllers"
)

func SetupRoutes(router *gin.Engine ) {
	router.GET("/getUserDetails", controllers.GetUserDetails)
}
