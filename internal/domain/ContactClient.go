package domain

import "gorm.io/gorm"

type ContactClient struct {
	gorm.Model
	ClientID  uint   `gorm:"not null;index"`
	Type      string `gorm:"type:varchar(50);not null"`
	Value     string `gorm:"type:text;not null"`
	IsActive  bool   `gorm:"default:true"`
	IsDefault bool   `gorm:"default:true"`
	IsPrimary bool   `gorm:"default:true"`
}
