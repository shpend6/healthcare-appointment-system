package models

import (
	"time"

	"gorm.io/gorm"
)

type Base struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
