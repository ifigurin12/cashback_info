package repo

import (
	"context"

	entity "cashback_info/domain/entities/card"

	"github.com/google/uuid"
)

type ICardRepo interface {
	ListCardsByUserID(ctx context.Context, userID uuid.UUID) ([]entity.Card, error)
}
