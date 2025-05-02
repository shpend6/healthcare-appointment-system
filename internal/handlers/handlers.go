package handlers

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"healthcare-appointment-system/internal/models"
	"healthcare-appointment-system/internal/models/dto"
	repositories "healthcare-appointment-system/internal/repository"
)

type Handler struct {
	PatientRepo     *repositories.PatientRepository
	AppointmentRepo *repositories.AppointmentRepository
}

// SetupRoutes registers all API routes
func SetupRoutes(r *gin.Engine, h *Handler) {
	// Patient endpoints
	patients := r.Group("/patients")
	{
		patients.GET("", h.GetPatients)
		patients.POST("", h.CreatePatient)
		patients.GET(":id", h.GetPatientByID)
		patients.GET(":id/appointments", h.GetAppointmentsByPatient)
	}

	// Appointment endpoints
	appts := r.Group("/appointments")
	{
		appts.GET("", h.GetAppointments)
		appts.POST("", h.CreateAppointment)
	}
}

// GetPatientByID returns a single patient by ID
func (h *Handler) GetPatientByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid patient ID"})
		return
	}

	patient, err := h.PatientRepo.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "patient not found"})
		return
	}
	c.JSON(http.StatusOK, patient)
}

// GetAppointments retrieves all appointments
func (h *Handler) GetAppointments(c *gin.Context) {
	limit, offset := getPaginationParams(c)
	appts, err := h.AppointmentRepo.GetAllPaginated(limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, appts)
}

// GetAppointmentsByPatient retrieves appointments for a specific patient
func (h *Handler) GetAppointmentsByPatient(c *gin.Context) {
	patientID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid patient ID"})
		return
	}

	appts, err := h.AppointmentRepo.GetByPatient(patientID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, appts)
}

// CreateAppointment handles creating a new appointment
func (h *Handler) CreateAppointment(c *gin.Context) {
	var dto dto.CreateAppointmentDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := h.PatientRepo.GetByID(dto.PatientID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "patient does not exist"})
		return
	}

	// Ensure date is in the future
	if dto.Date.Before(time.Now()) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "date must be in the future"})
		return
	}

	appt := models.Appointment{
		Base: models.Base{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Reason:    dto.Reason,
		Date:      dto.Date,
		PatientID: dto.PatientID,
	}
	if err := h.AppointmentRepo.Create(&appt); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, appt)
}

func (h *Handler) GetPatients(c *gin.Context) {
	limit, offset := getPaginationParams(c)
	patients, err := h.PatientRepo.GetAllPaginated(limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, patients)
}

func (h *Handler) CreatePatient(c *gin.Context) {
	var dto dto.CreatePatientDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if dto.DateOfBirth.After(time.Now()) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "date of birth must be in the past"})
		return
	}

	if strings.ToLower(dto.Gender) != "male" && strings.ToLower(dto.Gender) != "female" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "gender must be either male or female"})
		return
	}

	patient := models.Patient{
		Base: models.Base{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		FirstName:   dto.FirstName,
		LastName:    dto.LastName,
		Gender:      strings.ToLower(dto.Gender),
		DateOfBirth: dto.DateOfBirth,
		PhoneNumber: dto.PhoneNumber,
		Email:       dto.Email,
	}

	h.PatientRepo.Create(&patient)
	c.JSON(200, patient)
}
