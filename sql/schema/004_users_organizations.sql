-- +goose Up
ALTER TABLE USERS
ADD COLUMN organization_id uuid UNIQUE NOT NULL;

-- +goose Down
ALTER TABLE USERS
DROP COLUMN organization_id;