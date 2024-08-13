package models

import (
	"time"

	"github.com/google/uuid"
)

type Store struct {
	StoreID   uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primary_key"`
	StoreName string    `gorm:"type:varchar(100);not null"`
	Location  string    `gorm:"type:varchar(100)"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	CreatedBy uuid.UUID `gorm:"type:uuid"`
	UpdatedBy uuid.UUID `gorm:"type:uuid"`
}

func (Store) TableName() string {
	return "stores"
}
