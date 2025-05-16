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
)
