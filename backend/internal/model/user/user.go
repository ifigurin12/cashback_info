package user

import (
	"time"

	"github.com/google/uuid"
)

type Token struct {
	Token          string    `json:"token" binding:"required"`
	ExpirationTime time.Time `json:"expiration_time" binding:"required"`
}

type User struct {
	ID       uuid.UUID `json:"id" binding:"required"`
	Login    string    `json:"login" binding:"required"`
	Email    string    `json:"email" binding:"required"`
	RoleType RoleType  `json:"role_type" binding:"required"`
	Phone    *string   `json:"phone,omitempty"`
}

type RoleType string

const (
	Default RoleType = "default"
	Admin   RoleType = "admin"
)

func GenerateRoleTypeFromString(value string) *RoleType {
	var result RoleType
	switch value {
	case "default":
		result = Default
	case "admin":
		result = Admin
	default:
		return nil
	}

	return &result
}
