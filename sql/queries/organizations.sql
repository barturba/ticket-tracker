-- name: CreateOrganization :one
INSERT INTO ORGANIZATIONS (id, created_at, updated_at, name, user_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetOrganizationByUserID :one
SELECT * FROM ORGANIZATIONS WHERE USER_ID = $1;

-- name: UpdateOrganizationByUserID :one
UPDATE ORGANIZATIONS
SET UPDATED_AT = $2, NAME = $3
WHERE USER_ID = $1
RETURNING *;