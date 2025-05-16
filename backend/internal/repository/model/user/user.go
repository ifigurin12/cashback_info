package user

import "time"

type RoleType string

const (
	Default RoleType = "default"
	Admin   RoleType = "admin"
)

type User struct {
	ID           string    `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Username     string    `gorm:"type:varchar(50);unique;not null" json:"username"`
	Email        string    `gorm:"type:varchar(254);unique;not null" json:"email"`
	Phone        *string   `gorm:"type:varchar(16);unique" json:"phone"`
	PasswordHash string    `gorm:"type:varchar(255);not null" json:"password_hash"`
	RoleType     RoleType  `gorm:"type:role_types;default:'default'" json:"role_type"`
	DateCreated  time.Time `gorm:"type:timestamp;default:now();not null" json:"date_created"`
}
