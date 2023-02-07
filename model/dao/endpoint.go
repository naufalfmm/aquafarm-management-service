package dao

import (
	"time"

	"gorm.io/gorm"
)

type (
	Endpoint struct {
		ID     uint64 `gorm:"primaryKey"`
		Method string `gorm:"not null"`
		Path   string `gorm:"not null"`

		CreatedAt time.Time      `gorm:"not null"`
		UpdatedAt time.Time      `gorm:"not null"`
		DeletedAt gorm.DeletedAt `gorm:"null"`
		CreatedBy string         `gorm:"not null"`
		UpdatedBy string         `gorm:"not null"`
		DeletedBy *string        `gorm:"null"`
	}

	Endpoints []Endpoint
)

func (Endpoint) TableName() string {
	return "endpoints"
}

func (endpoints Endpoints) FindByMethodPath(method, path string) Endpoint {
	for _, endpoint := range endpoints {
		if endpoint.Method == method && endpoint.Path == path {
			return endpoint
		}
	}

	return Endpoint{}
}
