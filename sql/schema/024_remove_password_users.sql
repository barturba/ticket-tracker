-- +goose Up

ALTER TABLE users DROP COLUMN IF EXISTS password;

-- +goose Down

ALTER TABLE USERS
ADD COLUMN password TEXT;

