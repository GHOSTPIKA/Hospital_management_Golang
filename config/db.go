package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DB is a global variable to store the MongoDB client instance
var DB *mongo.Client

// init() ensures ConnectDB() runs automatically
func init() {
	ConnectDB()
}

// ConnectDB initializes the MongoDB connection
func ConnectDB() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("❌ Error connecting to MongoDB:", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("❌ Could not connect to MongoDB:", err)
	}

	fmt.Println("✅ Successfully connected to MongoDB!")
	DB = client
}

// GetCollection ensures DB is initialized before returning a collection
func GetCollection(collectionName string) *mongo.Collection {
	if DB == nil {
		log.Fatal("❌ Database connection is not initialized. Call ConnectDB() first.")
	}
	return DB.Database("hospital_management").Collection(collectionName)
}
