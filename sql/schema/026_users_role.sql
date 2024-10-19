-- +goose Up

ALTER TABLE users ADD role VARCHAR(255) NOT NULL DEFAULT 'user'; 

-- +goose Down

ALTER TABLE users
DROP COLUMN role;
