package cashback

import (
	model "cashback_info/internal/repository/model/category/cashback"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CategoryCashbackRepository interface {
	Create(items []model.CategoryCashbackDB) error
	List(cardIDs []uuid.UUID) ([]model.CategoryCashbackDB, error)
	Update(categoryCashback model.CategoryCashbackDB) error
	Delete(cardID uuid.UUID) error
}

type categoryCashbackRepository struct {
	db *gorm.DB
}

func NewCategoryCashbackRepository(db *gorm.DB) CategoryCashbackRepository {
	return &categoryCashbackRepository{db: db}
}

func (r *categoryCashbackRepository) Create(items []model.CategoryCashbackDB) error {
	if err := r.db.Create(&items).Error; err != nil {
		return err
	}
	return nil
}

func (r *categoryCashbackRepository) List(cardIDs []uuid.UUID) ([]model.CategoryCashbackDB, error) {
	var categoryCashbacks []model.CategoryCashbackDB
	if err := r.db.Preload("Category").Where("card_id IN ?", cardIDs).Find(&categoryCashbacks).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return categoryCashbacks, nil
}

func (r *categoryCashbackRepository) Update(categoryCashback model.CategoryCashbackDB) error {
	if err := r.db.Save(categoryCashback).Error; err != nil {
		return err
	}
	return nil
}

func (r *categoryCashbackRepository) Delete(cardID uuid.UUID) error {
	if err := r.db.Where("card_id = ?", cardID).Delete(&model.CategoryCashbackDB{}).Error; err != nil {
		return err
	}
	return nil
}
