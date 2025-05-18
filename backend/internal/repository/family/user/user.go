package user

import (
	model "cashback_info/internal/repository/model/family/user"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FamilyUserRepository interface {
	Create(familyUser model.FamilyUserDB) error
	GetByFamilyID(familyID uuid.UUID) ([]model.FamilyUserDB, error)
	GetByUserID(userID uuid.UUID) (*model.FamilyUserDB, error)
	Delete(familyID uuid.UUID, userID uuid.UUID) error
}

type familyUserRepository struct {
	db *gorm.DB
}

func NewFamilyUserRepository(db *gorm.DB) FamilyUserRepository {
	return &familyUserRepository{db: db}
}

func (f *familyUserRepository) Create(familyUser model.FamilyUserDB) error {
	if err := f.db.Create(&familyUser).Error; err != nil {
		return err
	}
	return nil
}

func (f *familyUserRepository) GetByFamilyID(familyID uuid.UUID) ([]model.FamilyUserDB, error) {
	var familyUsers []model.FamilyUserDB
	if err := f.db.Where("family_id = ?", familyID).Find(&familyUsers).Error; err != nil {
		return nil, err
	}
	return familyUsers, nil
}

func (f *familyUserRepository) GetByUserID(userID uuid.UUID) (*model.FamilyUserDB, error) {
	var familyUser model.FamilyUserDB
	if err := f.db.Where("user_id = ?", userID).Find(&familyUser).Error; err != nil {
		return nil, err
	}
	return &familyUser, nil
}

func (f *familyUserRepository) Delete(familyID uuid.UUID, userID uuid.UUID) error {
	if err := f.db.Where("family_id = ?", familyID).Where("user_id = ?", userID).Delete(&model.FamilyUserDB{}).Error; err != nil {
		return err
	}
	return nil
}
