package main
// main.go - Entry point for the MyAPI application
import (
	"github.com/Hyatus/myapi/config"
	"github.com/Hyatus/myapi/middleware" // Importing the middleware package for error handling
	"github.com/Hyatus/myapi/routes"
	"github.com/Hyatus/myapi/utils"
	"github.com/gin-gonic/gin" // Importing gin for HTTP routing
)

func main() {
	utils.InitLogger() 
	// Conectar con la DB
	cfg := config.LoadConfig() // Loading the configuration
	config.ConnectDB(cfg) // Connecting to the database using the connection string from the config

	r := gin.Default()
	r.Use(middleware.ErrorHandler()) // Using the error handler middleware
	routes.RegisterRoutes(r)
	utils.Log.Info("Starting server on port 8080") // Logging the server start
	r.Run(":8080") // Start the server on port 8080
}