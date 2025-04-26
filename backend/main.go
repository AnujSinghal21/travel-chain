package main
import (
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
	"github.com/anujSinghal21/travel-chain/api-interface/routes"
)

func main(){
	router := gin.Default()
	router.Use(cors.Default())
	routes.SetupRoutes(router)
	router.Run(":8080")	
}