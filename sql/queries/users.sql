-- name: CreateUser :one
INSERT INTO USERS (id, created_at, updated_at, name)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetUsers :many
SELECT * FROM users 
ORDER BY users.name ASC;

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