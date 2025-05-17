package category

import (
	entity "cashback_info/internal/repository/model/category"
	"errors"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(categories []entity.CategoryDB) error
	List(bankID int32) ([]entity.CategoryDB, error)
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) Create(categories []entity.CategoryDB) error {
	if err := r.db.Create(categories).Error; err != nil {
		return err
	}
	return nil
}

func (r *categoryRepository) List(bankID int32) ([]entity.CategoryDB, error) {
	var categories []entity.CategoryDB

	if err := r.db.Where("bank_id = ?", bankID).Find(&categories).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return categories, nil
}
