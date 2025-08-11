package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/Hyatus/myapi/controllers"
)

func RegisterRoutes(router *gin.Engine) {
	userRoutes := router.Group("/users")
	{
		userRoutes.GET("/", controllers.GetUsers) // Get all users
		userRoutes.GET("/:id", controllers.GetUserByID) // Get user by ID
		userRoutes.POST("/", controllers.CreateUser) // Create a new user
		userRoutes.PUT("/:id", controllers.UpdateUser) // Update user by ID
		userRoutes.DELETE("/:id", controllers.DeleteUser) // Delete user by ID
	}
			// Handle Wrong HTTP paths and write to log 
	router.NoRoute(func(c *gin.Context){
			controllers.HandleWrongPath(c, c.Request.URL.Path)
	})

}