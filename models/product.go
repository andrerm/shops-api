package models

import "time"

type Product struct {
	ProductID string    `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name      string    `gorm:"type:varchar(100);not null"`
	Price     float64   `gorm:"type:numeric;not null"`
	Stock     int       `gorm:"type:int;not null"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	CreatedBy string    `gorm:"type:uuid"`
	UpdatedBy string    `gorm:"type:uuid"`
	StoreID   string    `gorm:"type:uuid;not null"`
}
