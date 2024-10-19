-- +goose Up
CREATE TABLE IF NOT EXISTS permissions (
    id uuid DEFAULT gen_random_uuid(),
    code TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now(),

    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS users_permissions (
    user_id uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    permission_id uuid NOT NULL REFERENCES permissions(id) ON DELETE CASCADE,

    PRIMARY KEY (user_id, permission_id)
);

INSERT INTO permissions (code) 
VALUES ('users:read'), ('users:write'), ('users:delete'),
('cis:read'), ('cis:write'), ('cis:delete'),
('companies:read'), ('companies:write'), ('companies:delete'),
('permissions:read'), ('permissions:write'), ('permissions:delete'),
('users_permissions:read'), ('users_permissions:write'), ('users_permissions:delete');

-- +goose Down
DROP TABLE IF EXISTS users_permissions;
DROP TABLE IF EXISTS permissions;
