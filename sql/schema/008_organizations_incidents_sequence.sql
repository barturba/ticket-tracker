-- +goose Up
ALTER TABLE ORGANIZATIONS 
ADD COLUMN incidents_sequence BIGINT NOT NULL DEFAULT 1;

UPDATE ORGANIZATIONS 
SET incidents_sequence = 1 where id = '99ae48fb-89e8-4f7b-bd4c-b947b8fe8187';

-- +goose Down
ALTER TABLE ORGANIZATIONS 
DROP COLUMN incidents_sequence;