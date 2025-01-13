package card

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CardServer struct {
	ctx          context.Context
	postgresPool *pgxpool.Pool
}

func New(ctx context.Context, postgresPool *pgxpool.Pool) *CardServer {
	return &CardServer{ctx, postgresPool}
}
