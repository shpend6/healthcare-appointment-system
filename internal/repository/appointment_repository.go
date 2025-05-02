package repositories

import (
	"healthcare-appointment-system/internal/database"
	"healthcare-appointment-system/internal/models"

	"gorm.io/gorm"
)

// AppointmentRepository encapsulates database operations for appointments
type AppointmentRepository struct {
	DB *gorm.DB
}

// NewAppointmentRepository creates a new instance of AppointmentRepository
func NewAppointmentRepository() *AppointmentRepository {
	return &AppointmentRepository{DB: database.DB}
}

// GetAll retrieves all appointments
func (r *AppointmentRepository) GetAllPaginated(limit, offset int) ([]models.Appointment, error) {
	var appts []models.Appointment
	if err := r.DB.Limit(limit).Offset(offset).Find(&appts).Error; err != nil {
		return nil, err
	}
	return appts, nil
}

// GetByPatient retrieves all appointments for a given patient ID
func (r *AppointmentRepository) GetByPatient(patientID int) ([]models.Appointment, error) {
	var appts []models.Appointment
	if err := r.DB.Where("patient_id = ?", patientID).Find(&appts).Error; err != nil {
		return nil, err
	}
	return appts, nil
}

// Create adds a new appointment record
func (r *AppointmentRepository) Create(a *models.Appointment) error {
	return r.DB.Create(a).Error
}
