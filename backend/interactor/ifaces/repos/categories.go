package repo

import (
	"context"

	entity "cashback_info/domain/entities/card"

	"github.com/google/uuid"
)

type ICategoryRepo interface {
	ListCategories(ctx context.Context) ([]entity.Category, error)
	ListCategoriesByCardIDs(ctx context.Context, cardIDs []uuid.UUID) ([]entity.CardCategory, error)
}
