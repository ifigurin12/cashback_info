package card

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CategoryServer struct {
	ctx          context.Context
	postgresPool *pgxpool.Pool
}

func New(ctx context.Context, postgresPool *pgxpool.Pool) *CategoryServer {
	return &CategoryServer{ctx, postgresPool}
}
