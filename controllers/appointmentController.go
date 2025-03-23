package controllers

import (
	"context"
	"net/http"
	"time"

	"hospital-management/config"
	"hospital-management/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Appointment Collection
var appointmentCollection = config.GetCollection("appointments")

// Book a new appointment
func BookAppointment(c *gin.Context) {
	var appointment models.Appointment

	if err := c.ShouldBindJSON(&appointment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	appointment.ID = primitive.NewObjectID()
	appointment.CreatedAt = time.Now()
	appointment.UpdatedAt = time.Now()
	appointment.Status = "pending"

	_, err := appointmentCollection.InsertOne(context.TODO(), appointment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error booking appointment"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Appointment booked successfully!"})
}

// Get all appointments
func GetAppointments(c *gin.Context) {
	cursor, err := appointmentCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving appointments"})
		return
	}

	var appointments []models.Appointment
	if err = cursor.All(context.TODO(), &appointments); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding appointments"})
		return
	}

	c.JSON(http.StatusOK, appointments)
}
