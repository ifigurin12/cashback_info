package family

import (
	model "cashback_info/internal/repository/model/family"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FamilyRepository interface {
	Create(family model.FamilyDB) (*uuid.UUID, error)
	GetByID(id uuid.UUID) (*model.FamilyDB, error)
	GetByLeaderID(leaderID uuid.UUID) (*model.FamilyDB, error)
	Update(family model.FamilyDB) error
	Delete(id uuid.UUID) error
}

type familyRepository struct {
	db *gorm.DB
}

func NewFamilyRepository(db *gorm.DB) FamilyRepository {
	return &familyRepository{db: db}
}

func (f *familyRepository) Create(family model.FamilyDB) (*uuid.UUID, error) {
	if err := f.db.Create(&family).Error; err != nil {
		return nil, err
	}
	return &family.ID, nil
}

func (f *familyRepository) GetByID(id uuid.UUID) (*model.FamilyDB, error) {
	var family model.FamilyDB
	if err := f.db.Preload("Members").Preload("Leader").Where("id = ?", id).First(&family).Error; err != nil {
		return nil, err
	}
	return &family, nil
}

func (f *familyRepository) GetByLeaderID(leaderID uuid.UUID) (*model.FamilyDB, error) {
	var family model.FamilyDB
	if err := f.db.Where("leader_id = ?", leaderID).First(&family).Error; err != nil {
		return nil, err
	}
	return &family, nil
}

func (f *familyRepository) Update(family model.FamilyDB) error {
	if err := f.db.Save(family).Error; err != nil {
		return err
	}
	return nil
}

func (f *familyRepository) Delete(id uuid.UUID) error {
	if err := f.db.Delete(&model.FamilyDB{}, id).Error; err != nil {
		return err
	}
	return nil
}
