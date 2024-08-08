package models

import "time"

type Transaction struct {
	TransactionID string    `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	UserID        string    `gorm:"type:uuid;not null"`
	StoreID       string    `gorm:"type:uuid;not null"`
	ProductID     string    `gorm:"type:uuid;not null"`
	Quantity      int       `gorm:"type:int;not null"`
	TotalPrice    float64   `gorm:"type:numeric;not null"`
	CreatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	CreatedBy     string    `gorm:"type:uuid"`
	UpdatedBy     string    `gorm:"type:uuid"`
}
