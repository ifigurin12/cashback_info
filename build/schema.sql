-- Сначала создаем типы ENUM
CREATE TYPE bank_types AS ENUM ('vtb', 'alfa', 'tinkoff', 'pochta', 'gazprom');
CREATE TYPE role_types AS ENUM ('default', 'family_member', 'family_leader', 'admin');

-- Создаем таблицу users, так как на нее ссылаются другие таблицы
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(254) NOT NULL UNIQUE,
    phone VARCHAR(16) UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    role_type role_types DEFAULT 'default',
    date_created TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW()
);

-- Теперь создаем таблицы, которые ссылаются на users
CREATE TABLE cards (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(50) NOT NULL,
    user_id UUID NOT NULL,
    bank_type bank_types NOT NULL, 
    date_created TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    last_updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES users(id) 
    ON UPDATE CASCADE 
    ON DELETE CASCADE
);

CREATE TABLE categories (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(50) NOT NULL,
    bank_type bank_types NOT NULL, 
    date_created TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    description TEXT
);

-- Таблица для связи cards и categories
CREATE TABLE cards_categories (
    card_id UUID NOT NULL,
    category_id UUID NOT NULL,
    PRIMARY KEY (card_id, category_id),
    FOREIGN KEY (card_id) REFERENCES cards(id) ON DELETE CASCADE,
    FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE CASCADE
);

-- Таблица для словаря MCC-кодов
CREATE TABLE mcc_dictionary (
    id SERIAL PRIMARY KEY,
    code VARCHAR(4) NOT NULL UNIQUE, 
    description VARCHAR(255)
);

-- Таблица для связи categories и MCC-кодов
CREATE TABLE categories_mcc_codes (
    category_id UUID NOT NULL,
    mcc_code VARCHAR(4) NOT NULL,
    PRIMARY KEY (category_id, mcc_code),
    FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE CASCADE,
    FOREIGN KEY (mcc_code) REFERENCES mcc_dictionary(code) ON DELETE CASCADE
);

-- Таблица для семей (families)
CREATE TABLE families (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(50) NOT NULL,
    leader_id UUID NOT NULL,
    date_created TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    FOREIGN KEY (leader_id) REFERENCES users(id) ON UPDATE CASCADE
);

-- Таблица для связи users и families
CREATE TABLE families_users (
    user_id UUID NOT NULL,
    family_id UUID NOT NULL,
    PRIMARY KEY (user_id, family_id),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (family_id) REFERENCES families(id) ON DELETE CASCADE
);
