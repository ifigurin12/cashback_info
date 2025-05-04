package cashback

import (
	model "cashback_info/internal/model/category"
	"errors"

	"gorm.io/gorm"
)

type CategoryCashbackRepository interface {
	Create(items []model.CategoryCashback) error
	ListByParams(cardID string, page int, pageSize int) ([]model.CategoryCashback, error)
	Update(categoryCashback *model.CategoryCashback) error
	Delete(cardID, categoryID string) error
}

type categoryCashbackRepository struct {
	db *gorm.DB
}

func NewCategoryCashbackRepository(db *gorm.DB) CategoryCashbackRepository {
	return &categoryCashbackRepository{db: db}
}

func (r *categoryCashbackRepository) Create(items []model.CategoryCashback) error {
	if err := r.db.Create(&items).Error; err != nil {
		return err
	}
	return nil
}

func (r *categoryCashbackRepository) ListByParams(cardID string, page int, pageSize int) ([]model.CategoryCashback, error) {
	var categoryCashbacks []model.CategoryCashback
	if err := r.db.Where("card_id = ?", cardID).Find(&categoryCashbacks).Limit(pageSize).Offset((page - 1) * pageSize).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return categoryCashbacks, nil
}

func (r *categoryCashbackRepository) Update(categoryCashback *model.CategoryCashback) error {
	if err := r.db.Save(categoryCashback).Error; err != nil {
		return err
	}
	return nil
}

func (r *categoryCashbackRepository) Delete(cardID, categoryID string) error {
	if err := r.db.Where("card_id = ? AND category_id = ?", cardID, categoryID).Delete(&model.CategoryCashback{}).Error; err != nil {
		return err
	}
	return nil
}
