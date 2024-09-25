-- name: CreateConfigurationItem :one
INSERT INTO configuration_items (id, created_at, updated_at, name, company_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetConfigurationItemByID :one
SELECT * FROM configuration_items
WHERE id = $1;

-- name: GetConfigurationItems :many
SELECT * FROM configuration_items;

-- name: GetConfigurationItemsByCompanyID :many
SELECT * FROM configuration_items
WHERE company_id = $1;

-- name: UpdateConfigurationItem :one
UPDATE configuration_items
SET name = $2,
updated_at = $3
WHERE id = $1
RETURNING *;