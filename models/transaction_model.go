package models

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	TransactionID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primary_key"`
	UserID        uuid.UUID `gorm:"type:uuid"`
	StoreID       uuid.UUID `gorm:"type:uuid"`
	ProductID     uuid.UUID `gorm:"type:uuid"`
	PaymentTypeID uuid.UUID `gorm:"type:uuid"`
	BillID        uuid.UUID `gorm:"type:uuid"`
	Quantity      int       `gorm:"type:int;not null"`
	TotalPrice    float64   `gorm:"type:numeric;not null"`
	Date          time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	CreatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	CreatedBy     uuid.UUID `gorm:"type:uuid"`
	UpdatedBy     uuid.UUID `gorm:"type:uuid"`

	// Relationships
	User        User        `gorm:"foreignKey:UserID;references:UserID"`
	Store       Store       `gorm:"foreignKey:StoreID;references:StoreID"`
	Product     Product     `gorm:"foreignKey:ProductID;references:ProductID"`
	PaymentType PaymentType `gorm:"foreignKey:PaymentTypeID;references:PaymentTypeID"`
}

func (Transaction) TableName() string {
	return "transactions"
}
