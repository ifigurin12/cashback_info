CREATE TABLE banks (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE
);

CREATE TABLE categories (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(50) NOT NULL,
    bank_id INT NOT NULL, -- Изменено на bank_id
    date_created TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    description TEXT,
    FOREIGN KEY (bank_id) REFERENCES banks (id) ON DELETE CASCADE -- Связь с таблицей banks
);

CREATE TABLE
    categories (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
        title VARCHAR(50) NOT NULL,
        bank_id bank_types NOT NULL,
        description TEXT,
    );

CREATE TABLE
    category_cashbacks (
        card_id UUID NOT NULL, -- Идентификатор карты
        category_id UUID NOT NULL, -- Идентификатор категории
        cashback_percentage DECIMAL(5, 1) CHECK (
            cashback_percentage > 0
            AND cashback_percentage <= 100
        ) NOT NULL, -- Процент кэшбека от 0 до 100 с 1 знаком после запятой
        start_date TIMESTAMP WITHOUT TIME ZONE, -- Дата начала действия кэшбека
        end_date TIMESTAMP WITHOUT TIME ZONE, -- Дата окончания действия кэшбека
        limit
            DECIMAL(10, 2), -- Лимит кэшбека (например, максимальная сумма кэшбека)
            FOREIGN KEY (category_id) REFERENCES categories (id) ON DELETE CASCADE, -- Внешний ключ для category_id
            FOREIGN KEY (card_id) REFERENCES cards (id) ON DELETE CASCADE, -- Внешний ключ для card_id
            PRIMARY KEY (card_id, category_id) -- Составной первичный ключ
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