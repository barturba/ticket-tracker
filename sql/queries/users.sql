-- name: ListUsers :many
-- Retrieves a paginated list of users with optional filtering and sorting.
-- Parameters:
--   query: Search term for filtering (optional)
--   order_by: Column name for sorting (id, created_at, updated_at, last_name, first_name)
--   order_dir: Sort direction (ASC or DESC)
--   limit: Maximum number of records to return
--   offset: Number of records to skip
WITH filtered_users AS (
  SELECT *
  FROM users
  WHERE (
    CASE WHEN @query IS NOT NULL THEN
      email ILIKE '%' || @query || '%' OR
      first_name ILIKE '%' || @query || '%' OR
      last_name ILIKE '%' || @query || '%'
    ELSE true END
  )
)
SELECT 
  count(*) OVER() as total_count,
  *
FROM filtered_users
ORDER BY
  CASE 
    WHEN @order_by::varchar = 'id' AND @order_dir::varchar = 'ASC' THEN id
  END ASC NULLS LAST,
  CASE 
    WHEN @order_by::varchar = 'id' AND @order_dir::varchar = 'DESC' THEN id
  END DESC NULLS LAST,
  CASE 
    WHEN @order_by::varchar = 'created_at' AND @order_dir::varchar = 'ASC' THEN created_at
  END ASC NULLS LAST,
  CASE 
    WHEN @order_by::varchar = 'created_at' AND @order_dir::varchar = 'DESC' THEN created_at
  END DESC NULLS LAST,
  CASE 
    WHEN @order_by::varchar = 'updated_at' AND @order_dir::varchar = 'ASC' THEN updated_at
  END ASC NULLS LAST,
  CASE 
    WHEN @order_by::varchar = 'updated_at' AND @order_dir::varchar = 'DESC' THEN updated_at
  END DESC NULLS LAST,
  CASE 
    WHEN @order_by::varchar = 'last_name' AND @order_dir::varchar = 'ASC' THEN last_name
  END ASC NULLS LAST,
  CASE 
    WHEN @order_by::varchar = 'last_name' AND @order_dir::varchar = 'DESC' THEN last_name
  END DESC NULLS LAST,
  CASE 
    WHEN @order_by::varchar = 'first_name' AND @order_dir::varchar = 'ASC' THEN first_name
  END ASC NULLS LAST,
  CASE 
    WHEN @order_by::varchar = 'first_name' AND @order_dir::varchar = 'DESC' THEN first_name
  END DESC NULLS LAST,
  id ASC  -- Default sort for stable pagination
LIMIT $1 
OFFSET $2;

-- name: CountUsers :one
SELECT count(*) FROM users 
WHERE (first_name ILIKE '%' || @query || '%' or @query is NULL)
OR (last_name ILIKE '%' || @query || '%' or @query is NULL);

-- name: CreateUser :one
INSERT INTO USERS (id, created_at, updated_at, first_name, last_name, email)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetUser :one
SELECT * FROM USERS WHERE id = $1;

-- name: ListRecentUsers :many
SELECT * FROM users 
ORDER BY users.updated_at DESC
LIMIT $1 OFFSET $2;

-- name: UpdateUser :one
UPDATE users 
SET updated_at = $2, 
first_name = $3,
last_name = $4,
email = $5
WHERE ID = $1
RETURNING *;

-- name: DeleteUser :one
DELETE FROM users 
WHERE id = $1
RETURNING *;