-- name: CreateCompany :one
INSERT INTO companies (id, created_at, updated_at, name)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetCompanies :many
SELECT * from companies;

-- name: GetCompanyByID :one
SELECT * from companies
WHERE id = $1;

-- name: GetCompaniesFiltered :many
SELECT * FROM companies
WHERE (name ILIKE '%' || @query || '%')
LIMIT $1 OFFSET $2;

-- name: GetCompaniesFilteredCount :one
SELECT count(*) FROM companies
WHERE (name ILIKE '%' || @query || '%');