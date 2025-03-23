package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User struct represents common user fields (patients, doctors, admins)
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name" json:"name" validate:"required"`
	Email     string             `bson:"email" json:"email" validate:"required,email"`
	Password  string             `bson:"password" json:"password" validate:"required"`
	Role      string             `bson:"role" json:"role" validate:"required,oneof=admin doctor patient"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}
