package user

import (
	model "cashback_info/internal/model/family/user"

	"gorm.io/gorm"
)

type FamilyUserRepository interface {
	Create(familyUsers []model.FamiliesUser) error
	ListByParams(familyID string) ([]model.FamiliesUser, error)
	Delete(familyID string) error
}

type familyUserRepository struct {
	db *gorm.DB
}

func (f *familyUserRepository) Create(familyUsers []model.FamiliesUser) error {
	if err := f.db.Create(&familyUsers).Error; err != nil {
		return err
	}
	return nil
}

func (f *familyUserRepository) ListByParams(familyID string) ([]model.FamiliesUser, error) {
	panic("unimplemented")
}

func (f *familyUserRepository) Delete(familyID string) error {
	panic("unimplemented")
}

func NewFamilyUserRepository(db *gorm.DB) FamilyUserRepository {
	return &familyUserRepository{db: db}
}
