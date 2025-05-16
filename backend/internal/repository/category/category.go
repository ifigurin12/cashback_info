package category

import (
	entity "cashback_info/internal/model/category"
	"errors"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(categories []entity.Category) error
	ListByParams(bankID *uint, userID *string, page int, pageSize int) ([]entity.Category, error)
	Update(category entity.Category) error
	Delete(id string) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) Create(categories []entity.Category) error {
	if err := r.db.Create(categories).Error; err != nil {
		return err
	}
	return nil
}

func (r *categoryRepository) ListByParams(bankID *uint, userID *string, page int, pageSize int) ([]entity.Category, error) {
	var categories []entity.Category

	if err := r.db.Where("bank_id = ?", bankID).Or("user_id = ?", userID).Find(&categories).Limit(pageSize).Offset((page - 1) * pageSize).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return categories, nil
}

func (r *categoryRepository) Update(category entity.Category) error {
	if err := r.db.Save(category).Error; err != nil {
		return err
	}
	return nil
}

func (r *categoryRepository) Delete(id string) error {
	if err := r.db.Delete(&entity.Category{}, id).Error; err != nil {
		return err
	}
	return nil
}
