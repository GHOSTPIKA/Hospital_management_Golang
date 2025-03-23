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

// Patient Collection
var patientCollection = config.GetCollection("patients")

// Create a new patient
func CreatePatient(c *gin.Context) {
	var patient models.Patient

	if err := c.ShouldBindJSON(&patient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	patient.ID = primitive.NewObjectID()
	patient.CreatedAt = time.Now()
	patient.UpdatedAt = time.Now()

	_, err := patientCollection.InsertOne(context.TODO(), patient)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating patient"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Patient created successfully!"})
}

// Get all patients
func GetPatients(c *gin.Context) {
	cursor, err := patientCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving patients"})
		return
	}

	var patients []models.Patient
	if err = cursor.All(context.TODO(), &patients); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding patients"})
		return
	}

	c.JSON(http.StatusOK, patients)
}
