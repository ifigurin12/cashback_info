-- name: ListCardsByUserID :many
SELECT
    c.id,
    c.title,
    c.bank_type,
    c.date_created,
    c.last_updated_at
FROM
    cards c
WHERE
    c.user_id = $1;