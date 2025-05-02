package models

import (
	"time"
)

type Appointment struct {
	Base
	Reason    string
	Date      time.Time
	PatientID int
}
