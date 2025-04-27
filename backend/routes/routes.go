package routes

import (
	"github.com/gin-gonic/gin"
	"backend/controllers"
	"backend/middleware"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/")
	{
		api.POST("/register", controllers.RegisterHandler)
		api.POST("/login", controllers.LoginHandler)
		api.POST("/logout", controllers.LogoutHandler)

		protected := api.Group("/")
		protected.Use(middleware.AuthMiddleware())
		{
			protected.GET("/user", controllers.GetUserDetailsHandler)
			protected.GET("/ticket", controllers.GetTicketByTID)
			protected.GET("/tickets", controllers.GetAllTickets)
			protected.POST("/ticket/create", controllers.CreateTickets)
			protected.POST("/ticket/delete", controllers.DeleteTickets)
			protected.POST("/ticket/cancel", controllers.CancelTickets)
			protected.POST("/ticket/book", controllers.BookTickets)
			protected.POST("/ticket/status", controllers.GetTicketStatus)
		}
	}
}
