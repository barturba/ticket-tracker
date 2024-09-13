-- +goose Up
CREATE TYPE state_enum AS ENUM ('New', 'In Progress', 'Assigned', 'On Hold', 'Resolved');

-- +goose Down
DROP TYPE state_enum;