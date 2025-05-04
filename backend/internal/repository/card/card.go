package card

import (
	model "cashback_info/internal/model/Card"
	"errors"

	"gorm.io/gorm"
)

type CardRepository interface {
	Create(card model.Card) error
	ListByParams(userID string, page int, pageSize int) ([]model.Card, error)
	Update(card model.Card) error
	Delete(id string) error
}

type cardRepository struct {
	db *gorm.DB
}

func NewCardRepository(db *gorm.DB) CardRepository {
	return &cardRepository{db: db}
}

// Create implements CardRepository.
func (r *cardRepository) Create(card model.Card) error {
	if err := r.db.Create(&card).Error; err != nil {
		return err
	}
	return nil
}

func (r *cardRepository) ListByParams(userID string, page int, pageSize int) ([]model.Card, error) {
	var cards []model.Card
	if err := r.db.Where("user_id = ?", userID).Find(&cards).Limit(pageSize).Offset((page - 1) * pageSize).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return cards, nil
}

func (r *cardRepository) Update(card model.Card) error {
	if err := r.db.Save(card).Error; err != nil {
		return err
	}
	return nil
}

func (r *cardRepository) Delete(id string) error {
	if err := r.db.Delete(&model.Card{}, id).Error; err != nil {
		return err
	}
	return nil
}
