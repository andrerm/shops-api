package models

import (
	"time"

	"github.com/google/uuid"
)

type UserRole struct {
	UserRoleID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	UserID     uuid.UUID `gorm:"type:uuid;not null"`
	RoleID     uuid.UUID `gorm:"type:uuid;not null"`
	StoreID    uuid.UUID `gorm:"type:uuid;not null"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	CreatedBy  uuid.UUID `gorm:"type:uuid"`
	UpdatedBy  uuid.UUID `gorm:"type:uuid"`

	User User `gorm:"foreignKey:UserID"`
	Role Role `gorm:"foreignKey:RoleID"`
}

func (UserRole) TableName() string {
	return "user_roles"
}
