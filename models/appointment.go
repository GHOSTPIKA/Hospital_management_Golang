package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Appointment struct represents a doctor's appointment
type Appointment struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	PatientID primitive.ObjectID `bson:"patient_id" json:"patient_id"`
	DoctorID  primitive.ObjectID `bson:"doctor_id" json:"doctor_id"`
	Date      time.Time          `bson:"date" json:"date"`
	Status    string             `bson:"status" json:"status" validate:"required,oneof=pending confirmed completed cancelled"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}
