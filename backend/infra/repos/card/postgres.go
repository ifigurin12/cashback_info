package card

import (
	"context"
	"fmt"

	entity "cashback_info/domain/entities/card"
	utility "cashback_info/infra/repos/private"
	"cashback_info/infra/repos/private/db"

	"github.com/google/uuid"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresRepo struct {
	postgresPool *pgxpool.Pool
}

func NewCardRepo(postgresPool *pgxpool.Pool) *PostgresRepo {
	return &PostgresRepo{postgresPool}
}

func (r *PostgresRepo) ListCardsByUserID(ctx context.Context, userID uuid.UUID) ([]entity.Card, error) {
	queries := db.New(r.postgresPool)

	items, err := queries.ListCardsByUserID(ctx, userID)

	if err != nil {
		return nil, utility.TransformError(err)
	}

	result := []entity.Card{}

	for _, item := range items {
		bankType := entity.CreateBankTypeFromString(string(item.BankType))

		if bankType == nil {
			return nil, fmt.Errorf("unknown bank type: %s", item.BankType)
		}

		result = append(result, entity.Card{
			ID:            item.ID,
			Title:         item.Title,
			BankType:      *bankType,
			DateCreated:   item.DateCreated.Time,
			LastUpdatedAt: item.LastUpdatedAt.Time,
		})
	}

	return result, nil
}
