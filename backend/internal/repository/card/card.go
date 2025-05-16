package card

import (
	entity "cashback_info/internal/repository/model/card"
	"errors"

	"gorm.io/gorm"
)

type CardRepository interface {
	Create(card entity.Card) error
	ListByParams(userID string, page int, pageSize int) ([]entity.Card, error)
	Update(card entity.Card) error
	Delete(id string) error
}

type cardRepository struct {
	db *gorm.DB
}

func NewCardRepository(db *gorm.DB) CardRepository {
	return &cardRepository{db: db}
}

func (r *cardRepository) Create(card entity.Card) error {
	if err := r.db.Create(&card).Error; err != nil {
		return err
	}
	return nil
}

func (r *cardRepository) ListByParams(userID string, page int, pageSize int) ([]entity.Card, error) {
	var cards []entity.Card
	if err := r.db.Where("user_id = ?", userID).Find(&cards).Limit(pageSize).Offset((page - 1) * pageSize).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return cards, nil
}

func (r *cardRepository) Update(card entity.Card) error {
	if err := r.db.Save(card).Error; err != nil {
		return err
	}
	return nil
}

func (r *cardRepository) Delete(id string) error {
	if err := r.db.Delete(&entity.Card{}, id).Error; err != nil {
		return err
	}
	return nil
}
