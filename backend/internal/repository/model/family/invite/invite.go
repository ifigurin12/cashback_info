package invite

import (
	"cashback_info/internal/model/user"
	"cashback_info/internal/repository/model/family"

	"github.com/google/uuid"
)

type FamilyInviteDB struct {
	ID       uuid.UUID       `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	FamilyID uuid.UUID       `gorm:"type:uuid;not null"`
	UserID   uuid.UUID       `gorm:"type:uuid;not null"`
	Family   family.FamilyDB `gorm:"foreignKey:family_id;references:ID;constraint:OnDelete:CASCADE;"`
	User     user.User       `gorm:"foreignKey:user_id;references:ID;constraint:OnDelete:CASCADE;"`
}

func (FamilyInviteDB) TableName() string {
	return "families_invites"
}
