-- name: ListUsers :many
SELECT count(*) OVER(), * FROM users 
WHERE (email ILIKE '%' || @query || '%' or @query is NULL)
OR (first_name ILIKE '%' || @query || '%' or @query is NULL)
OR (last_name ILIKE '%' || @query || '%' or @query is NULL)
ORDER BY
CASE WHEN (@order_by::varchar = 'id' AND @order_dir::varchar = 'ASC') THEN id END ASC,
CASE WHEN (@order_by::varchar = 'id' AND @order_dir::varchar = 'DESC') THEN id END DESC,
CASE WHEN (@order_by::varchar = 'created_at' AND @order_dir::varchar = 'ASC') THEN created_at END ASC,
CASE WHEN (@order_by::varchar = 'created_at' AND @order_dir::varchar = 'DESC') THEN created_at END DESC,
CASE WHEN (@order_by::varchar = 'updated_at' AND @order_dir::varchar = 'ASC') THEN updated_at END ASC,
CASE WHEN (@order_by::varchar = 'updated_at' AND @order_dir::varchar = 'DESC') THEN updated_at END DESC,
CASE WHEN (@order_by::varchar = 'last_name' AND @order_dir::varchar = 'ASC') THEN last_name END ASC,
CASE WHEN (@order_by::varchar = 'last_name' AND @order_dir::varchar = 'DESC') THEN last_name END DESC,
CASE WHEN (@order_by::varchar = 'first_name' AND @order_dir::varchar = 'ASC') THEN first_name END ASC,
CASE WHEN (@order_by::varchar = 'first_name' AND @order_dir::varchar = 'DESC') THEN first_name END DESC,
id ASC 
LIMIT $1 OFFSET $2;

-- name: CountUsers :one
SELECT count(*) FROM users 
WHERE (first_name ILIKE '%' || @query || '%' or @query is NULL)
OR (last_name ILIKE '%' || @query || '%' or @query is NULL);

-- name: CreateUser :one
INSERT INTO USERS (id, created_at, updated_at, first_name, last_name, email)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetUser :one
SELECT * FROM USERS WHERE id = $1;

-- name: ListRecentUsers :many
SELECT * FROM users 
ORDER BY users.updated_at DESC
LIMIT $1 OFFSET $2;

-- name: UpdateUser :one
UPDATE users 
SET updated_at = $2, 
first_name = $3,
last_name = $4,
email = $5
WHERE ID = $1
RETURNING *;

-- name: DeleteUser :one
DELETE FROM users 
WHERE id = $1
RETURNING *;