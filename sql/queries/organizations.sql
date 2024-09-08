-- name: CreateOrganization :one
INSERT INTO ORGANIZATIONS (id, created_at, updated_at, name)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetOrganizationByID :one
SELECT * FROM ORGANIZATIONS WHERE ID = $1;