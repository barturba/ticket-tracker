-- name: ListIncidents :many
SELECT count(*) OVER(), * FROM incidents
LEFT JOIN users
ON incidents.assigned_to = users.id
WHERE (incidents.short_description ILIKE '%' || @query || '%' or @query is NULL)
OR (incidents.description ILIKE '%' || @query || '%' or @query is NULL)
OR (incidents.id::text ILIKE '%' || @query || '%' or @query is NULL)
ORDER BY
CASE WHEN (@order_by::varchar = 'created_at' AND @order_dir::varchar = 'ASC') THEN incidents.created_at END ASC,
CASE WHEN (@order_by::varchar = 'created_at' AND @order_dir::varchar = 'DESC') THEN incidents.created_at END DESC,
CASE WHEN (@order_by::varchar = 'updated_at' AND @order_dir::varchar = 'ASC') THEN incidents.updated_at END ASC,
CASE WHEN (@order_by::varchar = 'updated_at' AND @order_dir::varchar = 'DESC') THEN incidents.updated_at END DESC,
CASE WHEN (@order_by::varchar = 'id' AND @order_dir::varchar = 'ASC') THEN incidents.id END ASC,
CASE WHEN (@order_by::varchar = 'id' AND @order_dir::varchar = 'DESC') THEN incidents.id END DESC,
CASE WHEN (@order_by::varchar = 'short_description' AND @order_dir::varchar = 'ASC') THEN incidents.short_description END ASC,
CASE WHEN (@order_by::varchar = 'short_description' AND @order_dir::varchar = 'DESC') THEN incidents.short_description END DESC,
CASE WHEN (@order_by::varchar = 'description' AND @order_dir::varchar = 'ASC') THEN incidents.description END ASC,
CASE WHEN (@order_by::varchar = 'description' AND @order_dir::varchar = 'DESC') THEN incidents.description END DESC,
CASE WHEN (@order_by::varchar = 'first_name' AND @order_dir::varchar = 'ASC') THEN first_name END ASC,
CASE WHEN (@order_by::varchar = 'first_name' AND @order_dir::varchar = 'DESC') THEN first_name END DESC,
CASE WHEN (@order_by::varchar = 'last_name' AND @order_dir::varchar = 'ASC') THEN last_name END ASC,
CASE WHEN (@order_by::varchar = 'last_name' AND @order_dir::varchar = 'DESC') THEN last_name END DESC,
incidents.id ASC 
LIMIT $1 OFFSET $2;

-- name: CountIncidents :one
SELECT count(*) FROM incidents
LEFT JOIN users
ON incidents.assigned_to = users.id
WHERE (short_description ILIKE '%' || @query || '%' or @query is NULL)
OR (description ILIKE '%' || @query || '%' or @query is NULL)
OR (incidents.id::text ILIKE '%' || @query || '%' or @query is NULL);

-- name: CreateIncident :one
INSERT INTO incidents (id, created_at, updated_at, short_description, description, state, configuration_item_id, company_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: GetIncident :one
SELECT * FROM incidents
LEFT JOIN users
ON incidents.assigned_to = users.id
WHERE incidents.id = $1;

-- name: ListRecentIncidents :many
SELECT * FROM incidents
LEFT JOIN users
ON incidents.assigned_to = users.id
ORDER BY incidents.updated_at DESC
LIMIT $1 OFFSET $2;

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

-- name: DeleteIncident :one
DELETE FROM incidents 
WHERE id = $1
RETURNING *;
