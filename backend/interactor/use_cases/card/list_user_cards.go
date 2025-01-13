package card

import (
	"context"

	entity "cashback_info/domain/entities/card"
	dto "cashback_info/interactor/dtos/card"
	repo "cashback_info/interactor/ifaces/repos"

	uuid "github.com/google/uuid"
)

type ListUserCardsUseCase struct {
	CardRepo     repo.ICardRepo
	CategoryRepo repo.ICategoryRepo
}

func (u *ListUserCardsUseCase) Execute(ctx context.Context, inputDTO dto.ListUserCardsInputDTO) (*dto.ListUserCardsOutputDTO, error) {
	cards, err := u.CardRepo.ListCardsByUserID(ctx, inputDTO.UserID)

	if err != nil {
		return nil, err
	}

	cardIDs := []uuid.UUID{}

	for _, card := range cards {
		cardIDs = append(cardIDs, card.ID)
	}

	categories, err := u.CategoryRepo.ListCategoriesByCardIDs(ctx, cardIDs)

	if err != nil {
		return nil, err
	}

	categoriesByCardID := make(map[uuid.UUID][]entity.Category)
	for _, category := range categories {
		categoriesByCardID[category.CardID] = append(categoriesByCardID[category.CardID], category.Category)
	}

	var result []entity.CardWithCategories
	for _, card := range cards {
		cardWithCategories := entity.CardWithCategories{
			Card:       card,
			Categories: categoriesByCardID[card.ID],
		}
		result = append(result, cardWithCategories)
	}

	return &dto.ListUserCardsOutputDTO{
		Items: result,
	}, nil

}
