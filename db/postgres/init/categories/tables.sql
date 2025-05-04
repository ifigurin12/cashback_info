CREATE TABLE
    banks (
        id SERIAL PRIMARY KEY,
        name VARCHAR(50) NOT NULL UNIQUE
    );

CREATE TABLE
    categories (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
        title VARCHAR(50) NOT NULL,
        user_id UUID,
        bank_id INTEGER,
        description TEXT,
        FOREIGN KEY (user_id) REFERENCES categories (id) ON DELETE CASCADE,
        CHECK (
            user_id IS NOT NULL
            OR bank_id IS NOT NULL
        )
    );

CREATE TABLE
    category_cashbacks (
        card_id UUID NOT NULL,
        category_id UUID NOT NULL,
        cashback_percentage DECIMAL(5, 1) CHECK (
            cashback_percentage > 0
            AND cashback_percentage <= 100
        ) NOT NULL,
        cashback_limit INT CHECK (cashback_limit > 0) NOT NULL,
        start_date TIMESTAMP WITHOUT TIME ZONE,
        end_date TIMESTAMP WITHOUT TIME ZONE,
        FOREIGN KEY (category_id) REFERENCES categories (id) ON DELETE CASCADE,
        FOREIGN KEY (card_id) REFERENCES cards (id) ON DELETE CASCADE,
        PRIMARY KEY (card_id, category_id)
    );

CREATE TABLE
    mcc_dictionary (
        id SERIAL PRIMARY KEY,
        code VARCHAR(4) NOT NULL UNIQUE,
        description VARCHAR(255)
    );

CREATE TABLE
    categories_mcc_codes (
        category_id UUID NOT NULL,
        mcc_code VARCHAR(4) NOT NULL,
        PRIMARY KEY (category_id, mcc_code),
        FOREIGN KEY (category_id) REFERENCES categories (id) ON DELETE CASCADE,
        FOREIGN KEY (mcc_code) REFERENCES mcc_dictionary (code) ON DELETE CASCADE
    );