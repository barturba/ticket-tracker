-- name: CreateConfigurationItem :one
INSERT INTO configuration_items (id, created_at, updated_at, name, organization_id, company_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetConfigurationItemByID :one
SELECT * FROM configuration_items
WHERE id = $1;

-- name: GetConfigurationItemsByOrganizationID :many
SELECT * FROM configuration_items
WHERE organization_id = $1;

-- name: GetConfigurationItemsByCompanyID :many
SELECT * FROM configuration_items
WHERE company_id = $1;

-- name: UpdateConfigurationItem :one
UPDATE configuration_items
SET name = $3,
updated_at = $4
WHERE id = $1 AND organization_id = $2
RETURNING *;