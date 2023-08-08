-- +goose Up

CREATE TABLE users (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT  NULL,
    contact_no TEXT NOT NULL UNIQUE,
    user_key TEXT NOT NULL,
    is_admin BOOLEAN NOT NULL,
    hashed_password TEXT NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_id ON users(id);
CREATE INDEX IF NOT EXISTS idx_created_at ON users(created_at);
CREATE INDEX IF NOT EXISTS idx_email ON users(user_key);

-- +goose Down

DROP TABLE users;
