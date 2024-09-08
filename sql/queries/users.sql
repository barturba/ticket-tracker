-- name: CreateUser :one
INSERT INTO USERS (id, created_at, updated_at, name, organization_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetUserByAPIKey :one
SELECT * FROM USERS WHERE APIKEY = $1;