-- +goose Up
CREATE TABLE companies (
    id uuid PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL,
    organization_id UUID NOT NULL
);

ALTER TABLE companies 
    ADD CONSTRAINT fk_organizations
    FOREIGN KEY (organization_id)
    REFERENCES organizations(id)
    ON DELETE CASCADE;

-- +goose Down
DROP TABLE companies;