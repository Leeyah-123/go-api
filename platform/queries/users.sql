-- name: GetUser :one
SELECT id, email, password FROM users
WHERE email = $1 LIMIT 1;

-- name: GetUserById :one
SELECT id, email FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUsers :many
SELECT * FROM users
ORDER BY id;

-- name: AddUser :one
INSERT INTO users (
  email, password
) VALUES (
  $1, $2
)
RETURNING *;