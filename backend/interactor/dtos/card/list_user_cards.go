package card

import (
	entity "cashback_info/domain/entities/card"

	"github.com/google/uuid"
)

type ListUserCardsInputDTO struct {
	UserID uuid.UUID
}

type ListUserCardsOutputDTO struct {
	Items []entity.CardWithCategories
}
