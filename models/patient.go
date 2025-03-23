package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Patient struct represents a patient in the system
type Patient struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID         primitive.ObjectID `bson:"user_id" json:"user_id"`
	Age            int                `bson:"age" json:"age" validate:"required,min=0"`
	Gender         string             `bson:"gender" json:"gender" validate:"required,oneof=male female other"`
	Phone          string             `bson:"phone" json:"phone" validate:"required"`
	Address        string             `bson:"address" json:"address"`
	MedicalHistory []string           `bson:"medical_history" json:"medical_history"`
	CreatedAt      time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt      time.Time          `bson:"updated_at" json:"updated_at"`
}
