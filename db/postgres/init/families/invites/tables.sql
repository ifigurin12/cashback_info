CREATE TABLE
    families_invites (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
        family_id UUID NOT NULL,
        user_id UUID NOT NULL,
        FOREIGN KEY (family_id) REFERENCES families (id) ON DELETE CASCADE,
        FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
    );