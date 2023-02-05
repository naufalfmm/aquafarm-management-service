package dao

import "time"

// create table `ponds` (
// 	`id` bigint unsigned not null auto_increment,
// 	`farm_id` bigint unsigned not null,
// 	`code` varchar(25) not null,
// 	`description` varchar(255) not null default '',
// 	`wide` float not null,
// 	`long` float not null,
// 	`depth` float not null,
// 	`created_at` datetime default current_timestamp,
// 	`updated_at` datetime default current_timestamp,
// 	`deleted_at` datetime default null,
// 	`created_by` varchar(255) not null,
// 	`updated_by` varchar(255) not null,
// 	`deleted_by` varchar(255) null,
// 	`deleted_unix` int not null default 0,

// 	primary key (`id`),
// 	unique key `uq_pond_code_farm` (`farm_id`, `code`, `deleted_unix`)
// );

type (
	Pond struct {
		ID          uint64 `gorm:"primaryKey"`
		FarmID      uint64 `gorm:"not null"`
		Code        string `gorm:"not null"`
		Description string `gorm:"not null"`

		Wide  float64 `gorm:"not null"`
		Long  float64 `gorm:"not null"`
		Depth float64 `gorm:"not null"`

		CreatedAt   time.Time  `gorm:"not null"`
		UpdatedAt   time.Time  `gorm:"not null"`
		DeletedAt   *time.Time `gorm:"null"`
		CreatedBy   string     `gorm:"not null"`
		UpdatedBy   string     `gorm:"not null"`
		DeletedBy   *string    `gorm:"null"`
		DeletedUnix int        `gorm:"not null"`

		Farm Farm `gorm:"foreignKey:ID;references:FarmID"`
	}

	Ponds []Pond
)

func (p Pond) TableName() string {
	return "ponds"
}
