package main

import (
	"hospital-management/config"
	"hospital-management/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to MongoDB
	config.ConnectDB()

	// Initialize Gin router
	r := gin.Default()

	// Setup routes
	routes.SetupRoutes(r)

	// Start server
	r.Run(":8080") // Runs on localhost:8080a
}
