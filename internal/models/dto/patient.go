package dto

import "time"

type CreatePatientDTO struct {
	FirstName   string    `json:"firstName" binding:"required"`
	LastName    string    `json:"lastName" binding:"required"`
	Gender      string    `json:"gender" binding:"required"`
	DateOfBirth time.Time `json:"dateOfBirth" binding:"required"`
	PhoneNumber string    `json:"phoneNumber" binding:"required"`
	Email       string    `json:"email" binding:"required,email"`
}
