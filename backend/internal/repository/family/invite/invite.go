package invite

import (
	"cashback_info/internal/repository/model/family/invite"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FamilyInviteRepository interface {
	Create(invite invite.FamilyInviteDB) error
	ListByFamilyID(familyID uuid.UUID) ([]invite.FamilyInviteDB, error)
	ListByUserID(userID uuid.UUID) ([]invite.FamilyInviteDB, error)
	DeleteByID(inviteID uuid.UUID) error
	DeleteByUserID(userID uuid.UUID) error
}

type familyInviteRepository struct {
	db *gorm.DB
}

func NewFamilyInviteRepository(db *gorm.DB) FamilyInviteRepository {
	return &familyInviteRepository{db: db}
}

func (f *familyInviteRepository) Create(invite invite.FamilyInviteDB) error {
	if err := f.db.Create(&invite).Error; err != nil {
		return err
	}
	return nil
}

func (f *familyInviteRepository) ListByFamilyID(familyID uuid.UUID) ([]invite.FamilyInviteDB, error) {
	var invites []invite.FamilyInviteDB
	if err := f.db.Where("family_id = ?", familyID).Find(&invites).Error; err != nil {
		return nil, err
	}
	return invites, nil
}

func (f *familyInviteRepository) ListByUserID(userID uuid.UUID) ([]invite.FamilyInviteDB, error) {
	var invites []invite.FamilyInviteDB
	if err := f.db.Where("user_id = ?", userID).Find(&invites).Error; err != nil {
		return nil, err
	}
	return invites, nil
}

func (f *familyInviteRepository) DeleteByID(inviteID uuid.UUID) error {
	if err := f.db.Delete(&invite.FamilyInviteDB{}, inviteID).Error; err != nil {
		return err
	}
	return nil
}

func (f *familyInviteRepository) DeleteByUserID(userID uuid.UUID) error {
	if err := f.db.Where("user_id = ?", userID).Delete(&invite.FamilyInviteDB{}).Error; err != nil {
		return err
	}
	return nil
}
