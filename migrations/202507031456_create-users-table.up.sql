CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    email VARCHAR UNIQUE NOT NULL,
    encrypted_password VARCHAR UNIQUE NOT NULL,
    created_at TIMESTAMPTZ,
    creator_id BIGINT,
    updated_at TIMESTAMPTZ,
    updater_id BIGINT,
    deleted_at TIMESTAMPTZ,
    deleter_id BIGINT
);

INSERT INTO users (
    email,
    encrypted_password,
    created_at,
    creator_id
) VALUES (
    'test@gmail.com',
    'testtest',
    NOW(),
    1
);