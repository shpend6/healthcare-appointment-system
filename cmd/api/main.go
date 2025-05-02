package main

import (
	"healthcare-appointment-system/internal/database"
	"healthcare-appointment-system/internal/handlers"
	repositories "healthcare-appointment-system/internal/repository"
	"healthcare-appointment-system/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()

	r := gin.Default()

	r.Use(middleware.RequestLogger())

	h := &handlers.Handler{
		PatientRepo:     repositories.NewPatientRepository(),
		AppointmentRepo: repositories.NewAppointmentRepository(),
	}

	handlers.SetupRoutes(r, h)

	r.Run()
}
