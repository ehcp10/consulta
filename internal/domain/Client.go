package domain

import (
	"gorm.io/gorm"
	"time"
)

type Client struct {
	gorm.Model
	FirstName         string             `gorm:"not null" json:"first_name"`
	LastName          string             `gorm:"not null" json:"last_name"`
	CPF               string             `gorm:"size:11;not null;unique" json:"cpf"`
	Age               uint8              `gorm:"not null" json:"age"`
	Occupation        string             `gorm:"not null" json:"occupation"`
	MaritalStatus     string             `gorm:"not null" json:"marital_status"`
	Gender            uint8              `gorm:"not null" json:"gender"`
	BirthDate         time.Time          `gorm:"not null" json:"birth_date"`
	Contacts          []ContactClient    `gorm:"foreignKey:ClientID" json:"contacts"`
	EmergencyContacts []EmergencyContact `gorm:"foreignKey:ClientID" json:"emergency_contact"`
	Appointments      []Appointment      `gorm:"foreignKey:ClientID" json:"appointments"`
}
