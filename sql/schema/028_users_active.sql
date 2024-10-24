-- +goose Up
ALTER TABLE users 
ADD COLUMN active BOOLEAN DEFAULT TRUE;

-- +goose Down
ALTER TABLE users
DROP COLUMN active; 
