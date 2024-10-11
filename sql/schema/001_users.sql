-- +goose Up
CREATE TABLE USERS (
    id uuid PRIMARY KEY,
    created_at timestamp not null,
    updated_at timestamp not null,
    first_name varchar(50),
    last_name varchar(50)
);

-- +goose Down
DROP TABLE USERS;