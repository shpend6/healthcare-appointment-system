package repositories

import (
	"healthcare-appointment-system/internal/database"
	"healthcare-appointment-system/internal/models"

	"gorm.io/gorm"
)

// PatientRepository encapsulates database operations for patients
type PatientRepository struct {
	DB *gorm.DB
}

// NewPatientRepository creates a new instance of PatientRepository
func NewPatientRepository() *PatientRepository {
	return &PatientRepository{DB: database.DB}
}

// GetAll retrieves all patients
func (r *PatientRepository) GetAllPaginated(limit, offset int) ([]models.Patient, error) {
	var patients []models.Patient
	if err := r.DB.Limit(limit).Offset(offset).Find(&patients).Error; err != nil {
		return nil, err
	}
	return patients, nil
}

// GetByID retrieves a patient by its ID
func (r *PatientRepository) GetByID(id int) (*models.Patient, error) {
	var patient models.Patient
	if err := r.DB.First(&patient, id).Error; err != nil {
		return nil, err
	}
	return &patient, nil
}

// Create adds a new patient record
func (r *PatientRepository) Create(p *models.Patient) error {
	return r.DB.Create(p).Error
}
