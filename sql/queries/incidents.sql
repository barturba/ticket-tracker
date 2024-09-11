-- name: CreateIncident :one
INSERT INTO incidents (id, created_at, updated_at, short_description, organization_id, configuration_item_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetIncidentsByOrganizationID :many
SELECT * FROM incidents
WHERE organization_id = $1;