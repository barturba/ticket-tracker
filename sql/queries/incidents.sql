-- name: CreateIncident :one
INSERT INTO incidents (id, created_at, updated_at, short_description, description, state, configuration_item_id, company_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: GetIncidents :many
SELECT * FROM incidents
LEFT JOIN users
ON incidents.assigned_to = users.id
ORDER BY incidents.updated_at DESC;

-- name: GetIncidentById :one
SELECT * FROM incidents
LEFT JOIN users
ON incidents.assigned_to = users.id
WHERE incidents.id = $1;

-- name: GetIncidentsFiltered :many
SELECT * FROM incidents
LEFT JOIN users
ON incidents.assigned_to = users.id
WHERE (short_description ILIKE '%' || @query || '%' or @query is NULL)
OR (description ILIKE '%' || @query || '%' or @query is NULL)
ORDER BY incidents.updated_at DESC
LIMIT $1 OFFSET $2;

-- name: GetIncidentsFilteredCount :one
SELECT count(*) FROM incidents
LEFT JOIN users
ON incidents.assigned_to = users.id
WHERE (short_description ILIKE '%' || @query || '%' or @query is NULL)
OR (description ILIKE '%' || @query || '%' or @query is NULL);

-- name: GetIncidentsLatest :many
SELECT * FROM incidents
LEFT JOIN users
ON incidents.assigned_to = users.id
ORDER BY incidents.updated_at DESC
LIMIT $1 OFFSET $2;

-- name: GetIncidentsAsc :many
SELECT * FROM incidents
LEFT JOIN users
ON incidents.assigned_to = users.id
ORDER BY $1 ASC, id ASC;

-- name: GetIncidentsDesc :many
SELECT * FROM incidents
LEFT JOIN users
ON incidents.assigned_to = users.id
ORDER BY $1 DESC, id ASC;

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
SET updated_at = $2, 
company_id = $3,
configuration_item_id = $4,
description = $5, 
short_description = $6, 
state = $7,
assigned_to = $8
WHERE ID = $1
RETURNING *;

-- name: DeleteIncidentByID :one
DELETE FROM incidents 
WHERE id = $1
RETURNING *;
