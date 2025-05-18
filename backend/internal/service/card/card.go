package card

import (
	"cashback_info/internal/model/bank"
	"cashback_info/internal/model/card"
	"cashback_info/internal/model/category"
	"cashback_info/internal/model/category/cashback"
	cardrepo "cashback_info/internal/repository/card"
	cashbackrepo "cashback_info/internal/repository/category/cashback"
	repocard "cashback_info/internal/repository/model/card"

	"github.com/google/uuid"
)

type CreateCardArgs struct {
	Title  string
	BankID int32
	UserID uuid.UUID
}

type CardService interface {
	ListCards(userID uuid.UUID) ([]card.Card, error)
	CreateCard(args CreateCardArgs) (*uuid.UUID, error)
	DeleteCard(cardID uuid.UUID) error
	IsCardOwnedByUserID(cardID, userID uuid.UUID) (*bool, error)
}

type cardService struct {
	cardRepo     cardrepo.CardRepository
	cashbackRepo cashbackrepo.CategoryCashbackRepository
}

func NewCardService(cardRepo cardrepo.CardRepository, cashbackRepo cashbackrepo.CategoryCashbackRepository) CardService {
	return &cardService{cardRepo: cardRepo, cashbackRepo: cashbackRepo}
}

func (c *cardService) CreateCard(args CreateCardArgs) (*uuid.UUID, error) {
	cardID, err := c.cardRepo.Create(repocard.CardDB{
		Title:  args.Title,
		UserID: args.UserID,
		BankID: &args.BankID,
	})

	if err != nil {
		return nil, err
	}

	return cardID, nil
}

func (c *cardService) ListCards(userID uuid.UUID) ([]card.Card, error) {
	cards, err := c.cardRepo.List(userID)
	if err != nil {
		return nil, err
	}

	cardsIDs := make([]uuid.UUID, len(cards))
	for i, item := range cards {
		cardsIDs[i] = item.ID
	}

	cashbackCategories, err := c.cashbackRepo.List(cardsIDs)
	if err != nil {
		return nil, err
	}

	cashbackCategoryMap := make(map[uuid.UUID][]category.CategoryWithCashback)
	for _, item := range cashbackCategories {
		cashbackCategoryMap[item.CardID] = append(cashbackCategoryMap[item.CardID], category.CategoryWithCashback{
			Category: category.Category{
				ID:          item.CategoryID,
				Title:       item.Category.Title,
				Description: item.Category.Description,
			},
			Cashback: cashback.Cashback{
				Percentage: item.CashbackPercentage,
				Limit:      item.CashbackLimit,
				StartDate:  item.StartDate,
				EndDate:    item.EndDate,
			},
		})
	}

	result := make([]card.Card, len(cards))
	for i, item := range cards {
		result[i] = card.Card{
			ID:            item.ID,
			Title:         item.Title,
			LastUpdatedAt: item.LastUpdatedAt,
			Bank: &bank.Bank{
				ID:   item.Bank.ID,
				Name: item.Bank.Name,
			},
			Categories: cashbackCategoryMap[item.ID],
		}
	}

	return result, nil
}

func (c *cardService) DeleteCard(cardID uuid.UUID) error {
	return c.cardRepo.Delete(cardID)
}

func (c *cardService) IsCardOwnedByUserID(cardID, userID uuid.UUID) (*bool, error) {
	cards, err := c.cardRepo.List(userID)
	if err != nil {
		return nil, err
	}
	var result bool
	for _, item := range cards {
		if item.ID == cardID {
			result = true
			break
		}
	}

	return &result, nil
}
