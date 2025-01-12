CREATE TYPE bank_types AS ENUM ('vtb', 'alfa', 'tinkoff', 'pochta', 'gazprom');

CREATE TABLE categories (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(50) NOT NULL,
    bank_type bank_types NOT NULL, 
    description TEXT, 
)

CREATE TABLE mcc_dictionary (
    id SERIAL PRIMARY KEY,
    code VARCHAR(4) NOT NULL UNIQUE, 
    description VARCHAR(255) 
);

CREATE TABLE categories_mcc_codes (
    category_id UUID NOT NULL,
    mcc_code VARCHAR(4) NOT NULL,
    PRIMARY KEY (category_id, mcc_code),
    FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE CASCADE,
    FOREIGN KEY (mcc_code) REFERENCES mcc_dictionary(code) ON DELETE CASCADE
);