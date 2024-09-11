-- +goose Up
ALTER TABLE configuration_items
ADD COLUMN company_id UUID NOT NULL;

ALTER TABLE configuration_items
ADD CONSTRAINT fk_companies
FOREIGN KEY (company_id)
REFERENCES companies(id)
ON DELETE CASCADE;

-- +goose Down
ALTER TABLE configuration_items
DROP COLUMN company_id;