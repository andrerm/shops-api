package models

import (
	"time"

	"github.com/google/uuid"
)

//	type Role struct {
//		RoleID    uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
//		RoleName  string    `gorm:"type:varchar(50);not null"`
//		CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
//		UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
//		CreatedBy uuid.UUID `gorm:"type:uuid"`
//		UpdatedBy uuid.UUID `gorm:"type:uuid"`
//	}
type Role struct {
	RoleID    uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	RoleName  string    `gorm:"type:varchar(50);not null"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	CreatedBy uuid.UUID `gorm:"type:uuid"`
	UpdatedBy uuid.UUID `gorm:"type:uuid"`

	Users []User `gorm:"many2many:user_roles;"`
}

func (Role) TableName() string {
	return "roles"
}
