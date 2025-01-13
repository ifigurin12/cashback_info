-- name: ListCategoriesByCardIDs :many 
SELECT
    cc.card_id,
    c.id,
    c.title,
    c.bank_type,
    c.description,
    c.date_created
FROM
    cards_categories cc
    JOIN categories c ON cc.category_id = c.id
WHERE
    cc.card_id = ANY(sqlc.arg('card_ids')::UUID[]);

-- name: ListCategories :many
SELECT
    id,
    title,
    bank_type,
    date_created,
    description
FROM
    categories;
