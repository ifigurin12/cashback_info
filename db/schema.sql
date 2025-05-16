BEGIN;

CREATE TYPE role_types AS ENUM ('default', 'admin');

CREATE TABLE
    users (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
        username VARCHAR(50) NOT NULL UNIQUE,
        email VARCHAR(254) NOT NULL UNIQUE,
        phone VARCHAR(16) UNIQUE,
        password_hash VARCHAR(255) NOT NULL,
        role_type role_types DEFAULT 'default',
        date_created TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW ()
    );

CREATE TABLE families (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(50) NOT NULL,
    leader_id UUID NOT NULL,
    date_created TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW (),
    FOREIGN KEY (leader_id) REFERENCES users(id) ON UPDATE CASCADE
);

CREATE TABLE families_users (
    user_id UUID NOT NULL,
    family_id UUID NOT NULL,
    PRIMARY KEY (user_id, family_id),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (family_id) REFERENCES families(id) ON DELETE CASCADE
);

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
    cards (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
        title VARCHAR(50) NOT NULL,
        user_id UUID NOT NULL,
        bank_id INT,
        date_created TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW (),
        last_updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW (),
        FOREIGN KEY (user_id) REFERENCES users (id) ON UPDATE CASCADE ON DELETE CASCADE,
        FOREIGN KEY (bank_id) REFERENCES banks (id) ON DELETE CASCADE
    );

CREATE TABLE
    cards_categories (
        card_id UUID NOT NULL,
        category_id UUID NOT NULL,
        PRIMARY KEY (card_id, category_id),
        FOREIGN KEY (card_id) REFERENCES cards (id) ON DELETE CASCADE,
        FOREIGN KEY (category_id) REFERENCES categories (id) ON DELETE CASCADE
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

COMMIT;