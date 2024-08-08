package models

import "time"

type Store struct {
	StoreID   string    `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	StoreName string    `gorm:"type:varchar(100);not null"`
	Location  string    `gorm:"type:varchar(100)"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	CreatedBy string    `gorm:"type:uuid"`
	UpdatedBy string    `gorm:"type:uuid"`
}
