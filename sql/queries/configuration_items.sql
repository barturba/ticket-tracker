-- name: CreateConfigurationItem :one
INSERT INTO configuration_items (id, created_at, updated_at, name, organization_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetConfigurationItemsByOrganizationID :many
SELECT * FROM configuration_items
WHERE organization_id = $1;