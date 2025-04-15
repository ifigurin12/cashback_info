
CREATE TYPE role_types AS ENUM ('default', 'admin');


CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(254) NOT NULL UNIQUE,
    phone VARCHAR(16) UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    role_type role_type DEFAULT 'default',
    date_created TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW (),
); 



