-- +goose Up
CREATE TABLE ORGANIZATIONS (
    id uuid PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name text not null
);

-- +goose Down
DROP TABLE ORGANIZATIONS;