-- name: CreateMember :one
INSERT INTO members (
  full_name,
  latitude,
  longitude,
  phone_number,
  status
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: ListMembers :many
SELECT * FROM members
ORDER BY id;

-- name: GetMember :one
SELECT * FROM members 
WHERE id = $1 LIMIT 1;

-- name: UpdateMember :one
UPDATE members 
    SET latitude = $2, longitude = $3, full_name = $4, status = $5
WHERE id = $1
RETURNING *;

-- name: UpdateMemberName :one
UPDATE members 
    SET full_name = $2, status = $3
WHERE id = $1
RETURNING *;

-- name: DeleteMember :exec
DELETE FROM members 
WHERE id = $1;

-- name: GetMemberByNumber :one
SELECT * FROM members
WHERE phone_number = $1 LIMIT 1;