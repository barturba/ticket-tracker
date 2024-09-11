-- +goose Up
ALTER TABLE incidents
ADD COLUMN company_id UUID NOT NULL;

ALTER TABLE incidents
ADD CONSTRAINT fk_companies
FOREIGN KEY (company_id)
REFERENCES companies(id)
ON DELETE CASCADE;

-- +goose Down
ALTER TABLE incidents
DROP COLUMN company_id;