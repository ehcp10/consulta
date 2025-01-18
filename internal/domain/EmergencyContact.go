package domain

import "gorm.io/gorm"

type EmergencyContact struct {
	gorm.Model
	ClientID     uint   `gorm:"not null;index"`
	Name         string `gorm:"type:text;not null"`
	Relationship string `gorm:"type:text;not null"`
	Type         string `gorm:"type:text;not null"`
	Value        string `gorm:"type:text;not null"`
}
