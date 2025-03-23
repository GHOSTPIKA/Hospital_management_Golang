package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Doctor struct represents a doctor in the system
type Doctor struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID     primitive.ObjectID `bson:"user_id" json:"user_id"`
	Speciality string             `bson:"speciality" json:"speciality" validate:"required"`
	Phone      string             `bson:"phone" json:"phone" validate:"required"`
	Experience int                `bson:"experience" json:"experience" validate:"required,min=0"`
	Available  bool               `bson:"available" json:"available"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at" json:"updated_at"`
}
