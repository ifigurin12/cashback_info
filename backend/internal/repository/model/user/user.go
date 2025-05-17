package user

import (
	"time"

	"github.com/google/uuid"
)

type RoleTypeDB string

const (
	Default RoleTypeDB = "default"
	Admin   RoleTypeDB = "admin"
)

type UserDB struct {
	ID           uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Login        string     `gorm:"type:varchar(50);unique;not null" json:"login"`
	Email        string     `gorm:"type:varchar(254);unique;not null" json:"email"`
	Phone        *string    `gorm:"type:varchar(16);unique" json:"phone"`
	PasswordHash string     `gorm:"type:varchar(255);not null" json:"password_hash"`
	RoleType     RoleTypeDB `gorm:"type:role_types;default:'default'" json:"role_type"`
	DateCreated  time.Time  `gorm:"type:timestamp;default:now();not null" json:"date_created"`
}

func (u *UserDB) TableName() string {
	return "users"
}
