-- name: CreateUser :one
INSERT INTO USERS (id, created_at, updated_at, first_name, last_name, apikey, email, password)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: GetUsers :many
SELECT * FROM users 
WHERE (first_name ILIKE '%' || @query || '%' or @query is NULL)
OR (last_name ILIKE '%' || @query || '%' or @query is NULL)
ORDER BY users.updated_at DESC
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