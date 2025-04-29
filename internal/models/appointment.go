package models

import (
	"time"

	"github.com/google/uuid"
)

type Appointment struct {
	Base
	Reason    string
	Date      time.Time
	DoctorID  uuid.UUID
	PatientID uuid.UUID
}
