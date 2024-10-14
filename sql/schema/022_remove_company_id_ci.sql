-- +goose Up

ALTER TABLE configuration_items DROP COLUMN IF EXISTS company_id;

-- +goose Down

ALTER TABLE configuration_items 
ADD COLUMN company UUID NOT NULL;

