-- +goose Up
ALTER TABLE incidents
ADD COLUMN state state_enum NOT NULL DEFAULT 'New';

ALTER TABLE incidents
ADD COLUMN assigned_to UUID;

ALTER TABLE incidents
ADD CONSTRAINT fk_assigned_to
FOREIGN KEY (assigned_to)
REFERENCES users(id);

-- +goose Down
ALTER TABLE incidents
DROP COLUMN state;

ALTER TABLE incidents
DROP COLUMN assigned_to;