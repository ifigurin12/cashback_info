package bank

import (
	entity "cashback_info/internal/model/bank"
	repository "cashback_info/internal/repository/bank"
)

type BankService interface {
	List() ([]entity.Bank, error)
}

type bankService struct {
	repo repository.BankRepository
}

func NewBankService(repo repository.BankRepository) BankService {
	return &bankService{repo: repo}
}

func (b *bankService) List() ([]entity.Bank, error) {
	items, err := b.repo.List()

	result := make([]entity.Bank, len(items))
	for i, item := range items {
		result[i] = entity.Bank{ID: item.ID, Name: item.Name}
	}
	return result, err

}
