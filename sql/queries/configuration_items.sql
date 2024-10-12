-- name: CreateConfigurationItems :one
INSERT INTO configuration_items (id, created_at, updated_at, name, company_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetConfigurationItems :many
SELECT * FROM configuration_items 
WHERE (name ILIKE '%' || @query || '%' or @query is NULL)
ORDER BY
CASE WHEN (@order_by::varchar = 'id' AND @order_dir::varchar = 'ASC') THEN id END ASC,
CASE WHEN (@order_by::varchar = 'id' AND @order_dir::varchar = 'DESC') THEN id END DESC,
CASE WHEN (@order_by::varchar = 'created_at' AND @order_dir::varchar = 'ASC') THEN created_at END ASC,
CASE WHEN (@order_by::varchar = 'created_at' AND @order_dir::varchar = 'DESC') THEN created_at END DESC,
CASE WHEN (@order_by::varchar = 'updated_at' AND @order_dir::varchar = 'ASC') THEN updated_at END ASC,
CASE WHEN (@order_by::varchar = 'updated_at' AND @order_dir::varchar = 'DESC') THEN updated_at END DESC,
CASE WHEN (@order_by::varchar = 'name' AND @order_dir::varchar = 'ASC') THEN name END ASC,
CASE WHEN (@order_by::varchar = 'name' AND @order_dir::varchar = 'DESC') THEN name END DESC,
id ASC 
LIMIT $1 OFFSET $2;

-- name: GetConfigurationItemsCount :one
SELECT count(*) FROM configuration_items 
WHERE (name ILIKE '%' || @query || '%' or @query is NULL);

-- name: GetConfigurationItemsLatest :many
SELECT * FROM configuration_items 
ORDER BY configuration_items.updated_at DESC
LIMIT $1 OFFSET $2;

-- name: GetConfigurationItemsByID :one
SELECT * FROM configuration_items
WHERE id = $1;


-- name: GetConfigurationItemsByCompanyID :many
SELECT * FROM configuration_items
WHERE company_id = $1;

-- name: UpdateConfigurationItems :one
UPDATE configuration_items
SET name = $2,
updated_at = $3
WHERE id = $1
RETURNING *;

-- name: DeleteConfigurationItems :one
DELETE FROM configuration_items 
WHERE id = $1
RETURNING *;
