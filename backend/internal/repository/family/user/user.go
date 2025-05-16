package user

import (
	model "cashback_info/internal/model/family/user"

	"gorm.io/gorm"
)

type FamilyUserRepository interface {
	Create(familyUser model.FamilyUser) error
	ListByParams(familyID string) ([]model.FamilyUser, error)
	Delete(familyID string) error
}

type familyUserRepository struct {
	db *gorm.DB
}

func (f *familyUserRepository) Create(familyUser model.FamilyUser) error {
	if err := f.db.Create(&familyUser).Error; err != nil {
		return err
	}
	return nil
}

func (f *familyUserRepository) ListByParams(familyID string) ([]model.FamilyUser, error) {
	panic("unimplemented")
}

func (f *familyUserRepository) Delete(familyID string) error {
	panic("unimplemented")
}

func NewFamilyUserRepository(db *gorm.DB) FamilyUserRepository {
	return &familyUserRepository{db: db}
}
