package family

import (
	model "cashback_info/internal/model/family"

	"gorm.io/gorm"
)

type FamilyRepository interface {
	Create(family model.Family) error
	GetByID(id string) (*model.Family, error)
	Update(family model.Family) error
	Delete(id string) error
}

type familyRepository struct {
	db *gorm.DB
}

func NewFamilyRepository(db *gorm.DB) FamilyRepository {
	return &familyRepository{db: db}
}

// Create implements FamilyRepository.
func (f *familyRepository) Create(family model.Family) error {
	if err := f.db.Create(&family).Error; err != nil {
		return err
	}
	return nil
}

func (f *familyRepository) GetByID(id string) (*model.Family, error) {
	var family model.Family
	if err := f.db.First(&family, id).Error; err != nil {
		return nil, err
	}
	return &family, nil
}

func (f *familyRepository) Update(family model.Family) error {
	if err := f.db.Save(family).Error; err != nil {
		return err
	}
	return nil
}

func (f *familyRepository) Delete(id string) error {
	if err := f.db.Delete(&model.Family{}, id).Error; err != nil {
		return err
	}
	return nil
}
