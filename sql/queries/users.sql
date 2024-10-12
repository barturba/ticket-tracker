-- name: CreateUser :one
INSERT INTO USERS (id, created_at, updated_at, first_name, last_name, apikey, email, password)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: GetUsers :many
SELECT * FROM users 
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

-- name: GetUsersCount :one
SELECT count(*) FROM users 
WHERE (first_name ILIKE '%' || @query || '%' or @query is NULL)
OR (last_name ILIKE '%' || @query || '%' or @query is NULL);

-- name: GetUsersLatest :many
SELECT * FROM users 
ORDER BY users.updated_at DESC
LIMIT $1 OFFSET $2;

-- name: GetUsersByCompany :many
SELECT * FROM users 
LEFT JOIN companies 
ON users.assigned_to = users.id
ORDER BY users.name ASC;

-- name: GetUserByAPIKey :one
SELECT * FROM USERS WHERE APIKEY = $1;

-- name: GetUserByEmail :one
SELECT * FROM USERS WHERE email = $1;

-- name: GetUserByID :one
SELECT * FROM USERS WHERE id = $1;

-- name: UpdateUser :one
UPDATE users 
SET updated_at = $2, 
first_name = $3,
last_name = $4,
apikey = $5, 
email = $6, 
password = $7
WHERE ID = $1
RETURNING *;

-- name: DeleteUserByID :one
DELETE FROM users 
WHERE id = $1
RETURNING *;