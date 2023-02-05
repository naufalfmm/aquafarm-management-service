package dao

import (
	"time"

	"gorm.io/gorm"
)

type Farm struct {
	ID          uint64 `gorm:"primaryKey"`
	Code        string `gorm:"not null"`
	Description string `gorm:"not null"`

	Address    string `gorm:"not null"`
	Village    string `gorm:"not null"`
	District   string `gorm:"not null"`
	City       string `gorm:"not null"`
	Province   string `gorm:"not null"`
	PostalCode string `gorm:"not null"`

	Latitude  *float64 `gorm:"null"`
	Longitude *float64 `gorm:"null"`

	CreatedAt   time.Time      `gorm:"not null"`
	UpdatedAt   time.Time      `gorm:"not null"`
	DeletedAt   gorm.DeletedAt `gorm:"null"`
	CreatedBy   string         `gorm:"not null"`
	UpdatedBy   string         `gorm:"not null"`
	DeletedBy   *string        `gorm:"null"`
	DeletedUnix int            `gorm:"not null"`

	Ponds Ponds `gorm:"foreignKey:FarmID;references:ID"`
}

func (f Farm) TableName() string {
	return "farms"
}
