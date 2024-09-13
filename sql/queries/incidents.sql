-- name: CreateIncident :one
INSERT INTO incidents (id, created_at, updated_at, short_description, description, state, organization_id, configuration_item_id, company_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING *;

-- name: GetIncidentsByOrganizationID :many
SELECT * FROM incidents
LEFT JOIN users
ON incidents.assigned_to = users.id
WHERE organization_id = $1;