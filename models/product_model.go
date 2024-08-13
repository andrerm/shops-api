package models

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ProductID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name      string    `gorm:"type:varchar(100);not null"`
	Price     float64   `gorm:"type:numeric;not null"`
	Stock     int       `gorm:"type:int;not null"`
	Category  string    `gorm:"type:varchar(50)"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	CreatedBy uuid.UUID `gorm:"type:uuid"`
	UpdatedBy uuid.UUID `gorm:"type:uuid"`
	StoreID   uuid.UUID `gorm:"type:uuid"`
}

func (Product) TableName() string {
	return "products"
}
