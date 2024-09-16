-- +goose Up
ALTER TABLE users
ADD COLUMN email TEXT,
ADD COLUMN password TEXT;

ALTER TABLE users
ADD UNIQUE (email);

UPDATE users
SET email = 'john@gmail.com',
password = '$2a$10$DT8ZcidUx85Sv8iMqiQspe6N3cu44pzeuSwHcjjuYNTq6.Tsp7CzS';

ALTER TABLE users
ALTER COLUMN email
SET NOT NULL;

-- +goose Down
ALTER TABLE users 
DROP COLUMN email, 
DROP COLUMN password;