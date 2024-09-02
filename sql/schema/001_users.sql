-- +goose Up
CREATE TABLE USERS (
    id uuid PRIMARY KEY,
    created_at timestamp not null,
    updated_at timestamp not null,
    name text not null
);

-- +goose Down
DROP TABLE USERS;