-- name: CreateUser :one
INSERT INTO users (
  full_name,
  latitude,
  longitude,
  phone_number,
  status
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id;

-- name: GetUser :one
SELECT * FROM users 
WHERE id = $1 LIMIT 1;

-- name: UpdateUser :one
UPDATE users 
    SET latitude = $2, longitude = $3, full_name = $4, status = $5
WHERE id = $1
RETURNING *;

-- name: UpdateUserName :one
UPDATE users 
    SET full_name = $2, status = $3
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users 
WHERE id = $1;

-- name: GetUserByNumber :one
SELECT * FROM users
WHERE phone_number = $1 LIMIT 1;