-- +goose Up

ALTER TABLE users DROP COLUMN IF EXISTS apikey;

-- +goose Down

ALTER TABLE 
ALTER TABLE USERS
ADD COLUMN APIKEY VARCHAR(64) UNIQUE NOT NULL
DEFAULT encode(sha256(random()::text::bytea), 'hex');

