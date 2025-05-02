package models

import "time"

type Patient struct {
	Base
	FirstName   string
	LastName    string
	Gender      string
	DateOfBirth time.Time
	PhoneNumber string
	Email       string
}
