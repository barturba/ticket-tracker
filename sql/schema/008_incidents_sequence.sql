-- +goose Up
CREATE TABLE incidents_sequence (
    id uuid PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    organization_id UUID NOT NULL,
    sequence BIGINT NOT NULL
);

-- +goose Down
DROP TABLE incidents_sequence;