-- name: ListCompanies :many
SELECT count(*) OVER(), * FROM companies
WHERE (name ILIKE '%' || @query || '%' or @query is NULL)
ORDER BY
CASE WHEN (@order_by::varchar = 'id' AND @order_dir::varchar = 'ASC') THEN id END ASC,
CASE WHEN (@order_by::varchar = 'id' AND @order_dir::varchar = 'DESC') THEN id END DESC,
CASE WHEN (@order_by::varchar = 'created_at' AND @order_dir::varchar = 'ASC') THEN created_at END ASC,
CASE WHEN (@order_by::varchar = 'created_at' AND @order_dir::varchar = 'DESC') THEN created_at END DESC,
CASE WHEN (@order_by::varchar = 'updated_at' AND @order_dir::varchar = 'ASC') THEN updated_at END ASC,
CASE WHEN (@order_by::varchar = 'updated_at' AND @order_dir::varchar = 'DESC') THEN updated_at END DESC,
CASE WHEN (@order_by::varchar = 'name' AND @order_dir::varchar = 'ASC') THEN name END ASC,
CASE WHEN (@order_by::varchar = 'name' AND @order_dir::varchar = 'DESC') THEN name END DESC,
id ASC 
LIMIT $1 OFFSET $2;

-- name: CountCompanies :one
SELECT count(*) FROM companies
WHERE (name ILIKE '%' || @query || '%' or @query is NULL);

-- name: CreateCompany :one
INSERT INTO companies (id, created_at, updated_at, name)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetCompany :one
SELECT * from companies
WHERE id = $1;

-- name: ListRecentCompanies :many
SELECT * FROM companies 
ORDER BY companies.updated_at DESC
LIMIT $1 OFFSET $2;

-- name: UpdateCompany :one
UPDATE companies 
SET updated_at = $2, 
name = $3
WHERE ID = $1
RETURNING *;

-- name: DeleteCompany :one
DELETE FROM companies 
WHERE id = $1
RETURNING *;