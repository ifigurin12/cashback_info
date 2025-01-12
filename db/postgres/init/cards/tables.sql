CREATE TABLE cards (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(50) NOT NULL,
    user_id UUID NOT NULL,
    date_created TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW (),
    last_updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW (),
    FOREIGN KEY (user_id) REFERENCES users(id) 
    ON UPDATE CASCADE 
    ON DELETE CASCADE
);

CREATE TABLE cards_categories (
    card_id UUID NOT NULL,
    category_id UUID NOT NULL,
    PRIMARY KEY (card_id, category_id),
    FOREIGN KEY (card_id) REFERENCES cards(id) ON DELETE CASCADE,
    FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE CASCADE 
)
