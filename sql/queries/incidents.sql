-- name: CreateIncident :one
INSERT INTO incidents (id, created_at, updated_at, short_description, description, state, configuration_item_id, company_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: GetIncidents :many
SELECT * FROM incidents
LEFT JOIN users
ON incidents.assigned_to = users.id
ORDER BY incidents.updated_at DESC;

-- name: GetIncidentsBySearchTerm :many
SELECT * FROM incidents
LEFT JOIN users
ON incidents.assigned_to = users.id
WHERE short_description like $1
ORDER BY incidents.updated_at DESC;

-- name: GetIncidentsBySearchTermLimitOffset :many
SELECT *, count(*) OVER() AS full_count 
FROM incidents
LEFT JOIN users
ON incidents.assigned_to = users.id
WHERE short_description like $1 or short_description is NULL
ORDER BY incidents.updated_at DESC
LIMIT $2 OFFSET $3;

-- name: GetIncidentByID :one
SELECT * FROM incidents WHERE id = $1;

-- name: UpdateIncident :one
UPDATE incidents
SET updated_at = $2, description = $3, short_description = $4, state = $5
WHERE ID = $1
RETURNING *;
