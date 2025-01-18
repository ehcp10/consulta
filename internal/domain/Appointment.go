package domain

import (
	"gorm.io/gorm"
	"time"
)

type Appointment struct {
	gorm.Model
	ClientID uint      `gorm:"not null;index"`
	Date     time.Time `gorm:"not null"`
	Status   string    `gorm:"type:varchar(50);not null;default:'agendado'"`
	Cost     int64     `gorm:"not null"`
}
