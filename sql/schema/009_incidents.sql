-- +goose Up
CREATE TABLE incidents (
    id uuid PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    short_description TEXT NOT NULL,
    description TEXT,
    organization_id UUID NOT NULL,
    configuration_item_id UUID NOT NULL
);

ALTER TABLE incidents
    ADD CONSTRAINT fk_organizations
    FOREIGN KEY (organization_id)
    REFERENCES organizations(id)
    ON DELETE CASCADE;

ALTER TABLE incidents
    ADD CONSTRAINT fk_configuration_items
    FOREIGN KEY (configuration_item_id)
    REFERENCES configuration_items(id);

-- +goose Down
DROP TABLE incidents;