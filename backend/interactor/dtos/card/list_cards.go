package card

import (
	entity "cashback_info/domain/entities/card"

	"github.com/google/uuid"
)

type ListCardsInputDTO struct {
	UserID uuid.UUID
}

type ListCardsOutputDTO struct {
	Items []entity.CardWithCategories
}
