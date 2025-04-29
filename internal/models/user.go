package models

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
	Email        string
	PasswordHash string
	Role         Role
}
