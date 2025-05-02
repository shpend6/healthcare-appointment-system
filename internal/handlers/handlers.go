package handlers

import (
	"healthcare-appointment-system/internal/database"
	"healthcare-appointment-system/internal/models"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/patients", GetPatients)
	r.POST("/patients", CreatePatient)
}

func GetPatients(c *gin.Context) {
	var patients []models.User
	database.DB.Find(&patients)
	c.JSON(200, patients)
}

func CreatePatient(c *gin.Context) {
	var patient models.User
	if err := c.ShouldBindJSON(&patient); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&patient)
	c.JSON(200, patient)
}
