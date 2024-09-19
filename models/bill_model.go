package models

import "time"

type Bill struct {
	BillID    string    `gorm:"type:uuid;default:gen_random_uuid();primary_key"`
	BillName  string    `gorm:"type:varchar(200);not null"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	CreatedBy string    `gorm:"type:uuid"`
	UpdatedBy string    `gorm:"type:uuid"`
}

func (Bill) TableName() string {
	return "bills"
}
