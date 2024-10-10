-- name: CreateCompany :one
INSERT INTO companies (id, created_at, updated_at, name)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetCompanies :many
SELECT * from companies;

-- name: UpdateCompany :one
UPDATE companies 
SET updated_at = $2, 
name = $3
WHERE ID = $1
RETURNING *;

-- name: GetCompanyByID :one
SELECT * from companies
WHERE id = $1;

-- name: GetCompaniesFiltered :many
SELECT * FROM companies
WHERE (name ILIKE '%' || @query || '%' or @query is NULL)
ORDER BY companies.updated_at DESC
LIMIT $1 OFFSET $2;

-- name: GetCompaniesFilteredCount :one
SELECT count(*) FROM companies
WHERE (name ILIKE '%' || @query || '%' or @query is NULL);

-- name: DeleteCompanyByID :one
DELETE FROM companies 
WHERE id = $1
RETURNING *;