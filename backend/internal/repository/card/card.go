package card

import (
	entity "cashback_info/internal/repository/model/card"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CardRepository interface {
	Create(card entity.CardDB) (*uuid.UUID, error)
	List(userID uuid.UUID) ([]entity.CardDB, error)
	Update(card entity.CardDB) error
	Delete(id uuid.UUID) error
}

type cardRepository struct {
	db *gorm.DB
}

func NewCardRepository(db *gorm.DB) CardRepository {
	return &cardRepository{db: db}
}

func (r *cardRepository) Create(card entity.CardDB) (*uuid.UUID, error) {
	if err := r.db.Create(&card).Error; err != nil {
		return nil, err
	}
	return &card.ID, nil
}

func (r *cardRepository) List(userID uuid.UUID) ([]entity.CardDB, error) {
	var cards []entity.CardDB
	if err := r.db.Preload("Bank").Where("user_id = ?", userID).Find(&cards).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return cards, nil
}

func (r *cardRepository) Update(card entity.CardDB) error {
	if err := r.db.Save(card).Error; err != nil {
		return err
	}
	return nil
}

func (r *cardRepository) Delete(id uuid.UUID) error {
	if err := r.db.Delete(&entity.CardDB{}, id).Error; err != nil {
		return err
	}
	return nil
}
