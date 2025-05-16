package bank

import (
	model "cashback_info/internal/repository/model/bank"

	"gorm.io/gorm"
)

type BankRepository interface {
	List() ([]model.Bank, error)
}

type bankRepository struct {
	db *gorm.DB
}

func NewBankRepository(db *gorm.DB) BankRepository {
	return &bankRepository{db: db}
}

func (b *bankRepository) List() ([]model.Bank, error) {
	var banks []model.Bank
	if err := b.db.Find(&banks).Error; err != nil {
		return nil, err
	}
	return banks, nil
}
