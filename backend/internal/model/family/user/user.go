package user

import "cashback_info/internal/model/user"

type FamiliesUser struct {
	FamilyID string    `gorm:"type:uuid;not null" json:"family_id"`
	UserID   string    `gorm:"type:uuid;not null" json:"user_id"`
	User     user.User `gorm:"foreignkey:user_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user"`
}

func (FamiliesUser) TableName() string {
	return "family_users"
}
