-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, user_name)
VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING *;

-- name: GetUser :one
SELECT * from users
WHERE user_name = $1;

-- name: DeleteAllUsers :exec
DELETE from users;

-- name: GetAllUsers :many
SELECT * from users;