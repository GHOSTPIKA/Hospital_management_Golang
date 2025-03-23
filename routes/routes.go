package routes

import (
	"hospital-management/controllers"
	"hospital-management/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRoutes initializes all the API routes
func SetupRoutes(router *gin.Engine) {
	// Apply CORS middleware
	router.Use(middleware.CORSMiddleware())

	// Authentication routes
	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/register", controllers.Register)
		authRoutes.POST("/login", controllers.Login)
	}

	// Protected routes (Require authentication)
	apiRoutes := router.Group("/api")
	apiRoutes.Use(middleware.AuthMiddleware()) // Apply authentication middleware
	{
		// Patient routes
		apiRoutes.POST("/patients", controllers.CreatePatient)
		apiRoutes.GET("/patients", controllers.GetPatients)

		// Doctor routes
		apiRoutes.POST("/doctors", controllers.CreateDoctor)
		apiRoutes.GET("/doctors", controllers.GetDoctors)

		// Appointment routes
		apiRoutes.POST("/appointments", controllers.BookAppointment)
		apiRoutes.GET("/appointments", controllers.GetAppointments)
	}
}
