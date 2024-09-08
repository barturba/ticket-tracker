-- +goose Up
ALTER TABLE USERS
ADD CONSTRAINT fk_organizations
FOREIGN KEY (organization_id)
REFERENCES organizations(id)
ON DELETE CASCADE;

-- +goose Down
ALTER TABLE USERS
DROP CONSTRAINT fk_organizations;