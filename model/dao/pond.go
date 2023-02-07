package dao

import (
	"time"

	"github.com/naufalfmm/aquafarm-management-service/utils/orm"
	"gorm.io/gorm"
)

type (
	Pond struct {
		ID          uint64 `gorm:"primaryKey"`
		FarmID      uint64 `gorm:"not null"`
		Code        string `gorm:"not null"`
		Description string `gorm:"not null"`

		Wide  float64 `gorm:"not null"`
		Long  float64 `gorm:"not null"`
		Depth float64 `gorm:"not null"`

		CreatedAt   time.Time      `gorm:"not null"`
		UpdatedAt   time.Time      `gorm:"not null"`
		DeletedAt   gorm.DeletedAt `gorm:"null"`
		CreatedBy   string         `gorm:"not null"`
		UpdatedBy   string         `gorm:"not null"`
		DeletedBy   *string        `gorm:"null"`
		DeletedUnix int            `gorm:"not null"`

		Farm Farm `gorm:"foreignKey:ID;references:FarmID"`
	}

	Ponds []Pond

	PondsPagingResponse struct {
		orm.BasePagingResponse
		Items Ponds
	}
)

func (Pond) TableName() string {
	return "ponds"
}
