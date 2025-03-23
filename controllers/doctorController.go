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

// Doctor Collection
var doctorCollection = config.GetCollection("doctors")

// Create a new doctor
func CreateDoctor(c *gin.Context) {
	var doctor models.Doctor

	if err := c.ShouldBindJSON(&doctor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	doctor.ID = primitive.NewObjectID()
	doctor.CreatedAt = time.Now()
	doctor.UpdatedAt = time.Now()

	_, err := doctorCollection.InsertOne(context.TODO(), doctor)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating doctor"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Doctor created successfully!"})
}

// Get all doctors
func GetDoctors(c *gin.Context) {
	cursor, err := doctorCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving doctors"})
		return
	}

	var doctors []models.Doctor
	if err = cursor.All(context.TODO(), &doctors); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding doctors"})
		return
	}

	c.JSON(http.StatusOK, doctors)
}
