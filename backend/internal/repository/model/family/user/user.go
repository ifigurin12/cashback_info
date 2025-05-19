package user

import (
	"cashback_info/internal/repository/model/family"
	"cashback_info/internal/repository/model/user"

	"github.com/google/uuid"
)

type FamilyUserDB struct {
	FamilyID uuid.UUID       `gorm:"type:uuid;not null" json:"family_id"`
	UserID   uuid.UUID       `gorm:"type:uuid;not null" json:"user_id"`
	User     user.UserDB     `gorm:"foreignkey:user_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user"`
	Family   family.FamilyDB `gorm:"foreignKey:family_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"family"`
}

func (FamilyUserDB) TableName() string {
	return "families_users"
}
