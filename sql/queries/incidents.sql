-- name: CreateIncident :one
INSERT INTO incidents (id, created_at, updated_at, short_description, description, state, organization_id, configuration_item_id, company_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING *;

-- name: GetIncidentsByOrganizationID :many
SELECT * FROM incidents
LEFT JOIN users
ON incidents.assigned_to = users.id
WHERE organization_id = $1
ORDER BY incidents.updated_at DESC;

-- name: GetIncidentByID :one
SELECt * FROM incidents WHERE id = $1;

-- name: UpdateIncident :one
UPDATE incidents
SET updated_at = $2, description = $3, short_description = $4
WHERE ID = $1
RETURNING *;