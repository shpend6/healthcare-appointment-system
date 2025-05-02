package dto

import "time"

type CreateAppointmentDTO struct {
	PatientID int       `json:"patient_id" binding:"required"`
	Reason    string    `json:"reason" binding:"required"`
	Date      time.Time `json:"date" binding:"required"`
}
