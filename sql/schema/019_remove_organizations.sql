-- +goose Up
ALTER TABLE incidents DROP COLUMN organization_id;
ALTER TABLE companies DROP COLUMN organization_id;
DROP TABLE organizations;

-- +goose Down
CREATE TABLE organizations (
    id uuid PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name text not null,
    user_id uuid UNIQUE NOT NULL
);

ALTER TABLE INCIDENTS 
ADD COLUMN organization_id UUID NOT NULL;

ALTER TABLE incidents
    ADD CONSTRAINT fk_organizations
    FOREIGN KEY (organization_id)
    REFERENCES organizations(id)
    ON DELETE CASCADE;

ALTER TABLE companies
    ADD COLUMN organization_id UUID NOT NULL;

ALTER TABLE companies 
    ADD CONSTRAINT fk_organizations
    FOREIGN KEY (organization_id)
    REFERENCES organizations(id)
    ON DELETE CASCADE;