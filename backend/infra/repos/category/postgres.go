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

func NewPostgresRepo(postgresPool *pgxpool.Pool) *PostgresRepo {
	return &PostgresRepo{postgresPool}
}

func (r *PostgresRepo) ListCategories(ctx context.Context, userID uuid.UUID) ([]entity.Category, error) {
	queries := db.New(r.postgresPool)

	items, err := queries.ListCategories(ctx)

	if err != nil {
		return nil, utility.TransformError(err)
	}

	result := []entity.Category{}

	for _, item := range items {

		bankType := entity.CreateBankTypeFromString(string(item.BankType))

		if bankType == nil {
			return nil, fmt.Errorf("unknown bank type: %s", item.BankType)
		}

		result = append(result, entity.Category{
			ID:          item.ID,
			Title:       item.Title,
			BankType:    *bankType,
			DateCreated: item.DateCreated.Time,
			MCCCodes:    []entity.MССCode{},
			Description: item.Description,
		})
	}

	return result, nil
}

func (r *PostgresRepo) ListCategoriesByCardIDs(ctx context.Context, cardIDs []uuid.UUID) ([]entity.CardCategory, error) {
	queries := db.New(r.postgresPool)

	items, err := queries.ListCategoriesByCardIDs(ctx, cardIDs)

	if err != nil {
		return nil, utility.TransformError(err)
	}

	result := []entity.CardCategory{}

	for _, item := range items {

		bankType := entity.CreateBankTypeFromString(string(item.BankType))

		if bankType == nil {
			return nil, fmt.Errorf("unknown bank type: %s", item.BankType)
		}

		result = append(result, entity.CardCategory{
			CardID: item.CardID,
			Category: entity.Category{
				ID:          item.ID,
				Title:       item.Title,
				BankType:    *bankType,
				DateCreated: item.DateCreated.Time,
				MCCCodes:    []entity.MССCode{},
				Description: item.Description,
			},
		})
	}

	return result, nil
}
