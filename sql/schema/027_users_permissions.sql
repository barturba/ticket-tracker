-- +goose Up
CREATE TABLE IF NOT EXISTS permissions (
    id uuid DEFAULT gen_random_uuid(),
    code TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now(),

    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS users_permissions (
    user_id uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE
    permission_id uuid NOT NULL,

);
-- +goose Down

ALTER TABLE users
DROP COLUMN role;
