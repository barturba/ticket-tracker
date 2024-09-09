-- +goose Up
CREATE TABLE configuration_items (
    id uuid PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL,
    organization_id UUID NOT NULL
);

-- +goose Down
DROP TABLE configuration_items;