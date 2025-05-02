package models

import "time"

type Role string

const (
	RoleDoctor  Role = "doctor"
	RoleNurse   Role = "nurse"
	RolePatient Role = "patient"
)

type User struct {
	Base
	FirstName    string
	LastName     string
	Gender       string
	DateOfBirth  time.Time
	PhoneNumber  string
	Email        string
	PasswordHash string
	Role         Role
}
