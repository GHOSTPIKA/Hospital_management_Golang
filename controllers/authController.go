package controllers

import (
	"context"
	"fmt"
	"hospital-management/config"
	"hospital-management/models"
	"hospital-management/utils"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// User Collection
var userCollection = config.GetCollection("users")

// Register a new user (Admin, Doctor, or Patient)
func Register(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}
	user.Password = hashedPassword
	user.ID = primitive.NewObjectID()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	_, err = userCollection.InsertOne(context.TODO(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully!"})
}

// Login user and generate JWT token
func Login(c *gin.Context) {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Normalize email
	credentials.Email = strings.TrimSpace(strings.ToLower(credentials.Email))

	fmt.Println("🔍 Checking login for:", credentials.Email)

	var user models.User
	err := userCollection.FindOne(context.TODO(), bson.M{"email": credentials.Email}).Decode(&user)
	if err != nil {
		fmt.Println("❌ User not found in database:", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	fmt.Println("✅ User found:", user.Email)
	fmt.Println("👉 Hashed password from DB:", user.Password)

	// Debug: Check password before verification
	fmt.Println("👉 Input password:", credentials.Password)

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)); err != nil {
		fmt.Println("❌ Password verification failed:", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	fmt.Println("✅ Password matched successfully!")

	// Generate JWT token
	token, err := utils.GenerateJWT(user.ID.Hex(), user.Role)
	if err != nil {
		fmt.Println("❌ Error generating token:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
		return
	}

	fmt.Println("✅ Login successful!")
	c.JSON(http.StatusOK, gin.H{"token": token})
}
