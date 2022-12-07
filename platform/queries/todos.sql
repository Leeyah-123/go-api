-- name: GetUserTodos :many
SELECT * FROM todos
WHERE user_id = $1;

-- name: GetTodoById :one
SELECT * FROM todos
WHERE id = $1 AND user_id = $2
LIMIT 1;

-- name: AddTodo :one
INSERT INTO todos (
  user_id, title, completed
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: UpdateTodoById :one
UPDATE todos
  SET completed = $2
WHERE id = $1
RETURNING *;

-- name: DeleteTodoById :exec
DELETE FROM todos
WHERE id = $1;