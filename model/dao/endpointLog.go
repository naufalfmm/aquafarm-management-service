package dao

import "time"

type (
	EndpointLog struct {
		ID                 uint64  `gorm:"primaryKey"`
		EndpointID         uint64  `gorm:"not null"`
		RequestID          string  `gorm:"not null"`
		Uri                string  `gorm:"not null"`
		Query              string  `gorm:"not null"`
		UserAgent          string  `gorm:"not null"`
		IpAddress          string  `gorm:"not null"`
		RequestedBy        *string `gorm:"null"`
		ResponseStatusCode int     `gorm:"not null"`
		StartAt            int64   `gorm:"not null"`
		EndAt              *int64  `gorm:"null"`

		CreatedAt time.Time `gorm:"not null"`
		UpdatedAt time.Time `gorm:"not null"`
		CreatedBy string    `gorm:"not null"`
		UpdatedBy string    `gorm:"not null"`

		Endpoint Endpoint `gorm:"->;foreignKey:ID;references:EndpointID"`
	}
)

func (EndpointLog) TableName() string {
	return "endpoint_logs"
}
