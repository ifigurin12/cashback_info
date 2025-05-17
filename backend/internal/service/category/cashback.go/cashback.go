package cashback

import (
	entity "cashback_info/internal/model/category/cashback"
	"cashback_info/internal/repository/category/cashback"
	repositoryentity "cashback_info/internal/repository/model/category/cashback"

	"github.com/google/uuid"
)

type CategoryCashbackService interface {
	Create(cashbackItems []entity.Cashback, categoryIDs []uuid.UUID, cardID uuid.UUID) error
}

type categoryCashbackService struct {
	repository cashback.CategoryCashbackRepository
}

func NewCategoryCashbackService(repository cashback.CategoryCashbackRepository) CategoryCashbackService {
	return &categoryCashbackService{repository: repository}
}

func (c *categoryCashbackService) Create(cashbackItems []entity.Cashback, categoryIDs []uuid.UUID, cardID uuid.UUID) error {
	itemsToCreate := make([]repositoryentity.CategoryCashbackDB, len(cashbackItems))

	for i, cashbackItem := range cashbackItems {

		itemsToCreate[i] = repositoryentity.CategoryCashbackDB{
			CardID:             cardID,
			CategoryID:         categoryIDs[i],
			CashbackPercentage: cashbackItem.Percentage,
			StartDate:          cashbackItem.StartDate,
			EndDate:            cashbackItem.EndDate,
			CashbackLimit:      cashbackItem.Limit,
		}
	}

	return c.repository.Create(itemsToCreate)
}
