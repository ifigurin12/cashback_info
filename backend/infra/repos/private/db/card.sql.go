// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: card.sql

package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const listCardsByUserID = `-- name: ListCardsByUserID :many
SELECT
    c.id,
    c.title,
    c.bank_type,
    c.date_created,
    c.last_updated_at
FROM
    cards c
WHERE
    c.user_id = $1
`

type ListCardsByUserIDRow struct {
	ID            uuid.UUID
	Title         string
	BankType      BankTypes
	DateCreated   pgtype.Timestamp
	LastUpdatedAt pgtype.Timestamp
}

func (q *Queries) ListCardsByUserID(ctx context.Context, userID uuid.UUID) ([]ListCardsByUserIDRow, error) {
	rows, err := q.db.Query(ctx, listCardsByUserID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListCardsByUserIDRow
	for rows.Next() {
		var i ListCardsByUserIDRow
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.BankType,
			&i.DateCreated,
			&i.LastUpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
