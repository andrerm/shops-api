package models

import (
	"time"

	"github.com/google/uuid"
)

//	type User struct {
//		UserID    uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
//		Name      string    `gorm:"type:varchar(100);not null"`
//		Email     string    `gorm:"type:varchar(100);unique;not null"`
//		Password  string    `gorm:"type:varchar(100);not null"`
//		CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
//		UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
//		CreatedBy uuid.UUID `gorm:"type:uuid"`
//		UpdatedBy uuid.UUID `gorm:"type:uuid"`
//	}
type User struct {
	UserID    uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name      string    `gorm:"type:varchar(100);not null"`
	Email     string    `gorm:"type:varchar(100);unique;not null"`
	Password  string    `gorm:"type:varchar(100);not null"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	CreatedBy uuid.UUID `gorm:"type:uuid"`
	UpdatedBy uuid.UUID `gorm:"type:uuid"`

	Roles []Role `gorm:"many2many:user_roles;"`
}

func (User) TableName() string {
	return "users"
}
