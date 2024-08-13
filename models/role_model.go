package models

import "time"

type Role struct {
	RoleID    string    `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	RoleName  string    `gorm:"type:varchar(50);not null"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	CreatedBy string    `gorm:"type:uuid"`
	UpdatedBy string    `gorm:"type:uuid"`
}

func (Role) TableName() string {
	return "roles"
}
