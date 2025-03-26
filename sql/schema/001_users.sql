-- +goose Up
CREATE TABLE users(
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    upated_at TIMESTAMP NOT NULL,
    name TEXT UNIQUE NOT NULL
);

-- +goose Down
DROP TABLE users;
