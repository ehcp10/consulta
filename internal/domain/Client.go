package domain

import (
	"gorm.io/gorm"
	"time"
)

type Client struct {
	gorm.Model
	FistName         string             `gorm:"varchar(255);not null"`
	LastName         string             `gorm:"varchar(255);not null"`
	CPF              string             `gorm:"varchar(11);not null"`
	age              uint8              `gorm:"not null"`
	occupation       string             `gorm:"varchar(255);not null"`
	MaritalStatus    string             `gorm:"varchar(255);not null"`
	Gender           uint8              `gorm:"not null"`
	BirthDate        time.Time          `gorm:"not null"`
	Contacts         []ContactClient    `gorm:"foreignkey:ClientID"`
	EmergencyContact []EmergencyContact `gorm:"foreignkey:ClientID"`
	Appointment      []Appointment      `gorm:"foreignkey:ClientID"`
}
