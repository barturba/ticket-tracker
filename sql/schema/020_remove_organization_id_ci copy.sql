-- +goose Up

ALTER TABLE configuration_items DROP COLUMN IF EXISTS organization_id;

-- +goose Down

ALTER TABLE configuration_items 
ADD COLUMN organization_id UUID NOT NULL;

ALTER TABLE configuration_items 
    ADD CONSTRAINT fk_organizations
    FOREIGN KEY (organization_id)
    REFERENCES organizations(id)
    ON DELETE CASCADE;
