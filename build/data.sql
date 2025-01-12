-- Вставляем 3 пользователей
INSERT INTO users (id, username, email, phone, password_hash, role_type, date_created)
VALUES 
    (gen_random_uuid(), 'user1', 'user1@example.com', '1234567890', 'hashed_password_1', 'default', NOW()),
    (gen_random_uuid(), 'user2', 'user2@example.com', '0987654321', 'hashed_password_2', 'family_member', NOW()),
    (gen_random_uuid(), 'user3', 'user3@example.com', '1112223333', 'hashed_password_3', 'family_leader', NOW());

-- Получаем ID пользователей для дальнейшего использования
WITH users_data AS (
    SELECT id FROM users WHERE username IN ('user1', 'user2', 'user3')
)
SELECT * FROM users_data;

-- Вставляем одну семью, где user3 является лидером
INSERT INTO families (id, title, leader_id, date_created)
VALUES (gen_random_uuid(), 'The Awesome Family', (SELECT id FROM users WHERE username = 'user3'), NOW());

-- Получаем ID семьи для дальнейшего использования
WITH family_data AS (
    SELECT id FROM families WHERE title = 'The Awesome Family'
)
SELECT * FROM family_data;

-- Связываем пользователей с семьей (user2 и user3)
INSERT INTO families_users (user_id, family_id)
VALUES
    ((SELECT id FROM users WHERE username = 'user2'), (SELECT id FROM families WHERE title = 'The Awesome Family')),
    ((SELECT id FROM users WHERE username = 'user3'), (SELECT id FROM families WHERE title = 'The Awesome Family'));

-- Вставляем по одной карте для каждого пользователя
INSERT INTO cards (id, title, user_id, bank_type,  date_created, last_updated_at)
VALUES
    (gen_random_uuid(), 'Card for user1', (SELECT id FROM users WHERE username = 'user1'), 'vtb', NOW(), NOW()),
    (gen_random_uuid(), 'Card for user2', (SELECT id FROM users WHERE username = 'user2'), 'alfa', NOW(), NOW()),
    (gen_random_uuid(), 'Card 1 for user3', (SELECT id FROM users WHERE username = 'user3'), 'vtb', NOW(), NOW()),
    (gen_random_uuid(), 'Card 2 for user3', (SELECT id FROM users WHERE username = 'user3'), 'alfa', NOW(), NOW()),
    (gen_random_uuid(), 'Card 3 for user3', (SELECT id FROM users WHERE username = 'user3'), 'alfa', NOW(), NOW());



-- Вставка категорий
INSERT INTO categories (id, title, bank_type, description)
VALUES 
    (gen_random_uuid(), 'Groceries', 'vtb', 'Everyday groceries and supermarkets'),
    (gen_random_uuid(), 'Travel', 'alfa', 'Flights, hotels, and travel-related expenses'),
    (gen_random_uuid(), 'Restaurants', 'vtb', 'Dining out and restaurants'),
    (gen_random_uuid(), 'Electronics', 'alfa', 'Gadgets, electronics, and tech stores');

-- Вставка MCC-кодов
INSERT INTO mcc_dictionary (code, description)
VALUES 
    ('5411', 'Grocery Stores and Supermarkets'),
    ('4511', 'Airlines and Air Carriers'),
    ('5812', 'Eating Places and Restaurants'),
    ('5732', 'Electronics Stores');

-- Привязка MCC-кодов к категориям
INSERT INTO categories_mcc_codes (category_id, mcc_code)
VALUES 
    ((SELECT id FROM categories WHERE title = 'Groceries'), '5411'),
    ((SELECT id FROM categories WHERE title = 'Travel'), '4511'),
    ((SELECT id FROM categories WHERE title = 'Restaurants'), '5812'),
    ((SELECT id FROM categories WHERE title = 'Electronics'), '5732');

-- Привязка категорий к картам (например, 1 категория для каждой карты)
INSERT INTO cards_categories (card_id, category_id)
VALUES 
    ((SELECT id FROM cards WHERE title = 'Card for user1'), (SELECT id FROM categories WHERE title = 'Groceries')),
    ((SELECT id FROM cards WHERE title = 'Card for user2'), (SELECT id FROM categories WHERE title = 'Travel')),
    ((SELECT id FROM cards WHERE title = 'Card 1 for user3'), (SELECT id FROM categories WHERE title = 'Restaurants')),
    ((SELECT id FROM cards WHERE title = 'Card 2 for user3'), (SELECT id FROM categories WHERE title = 'Electronics')),
    ((SELECT id FROM cards WHERE title = 'Card 3 for user3'), (SELECT id FROM categories WHERE title = 'Travel'));
