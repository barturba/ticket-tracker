-- name: CreateCompany :one
INSERT INTO companies (id, created_at, updated_at, name, organization_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetCompaniesByOrganizationID :many
SELECT * from companies
WHERE organization_id = $1;