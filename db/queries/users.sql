-- name: CreateUser :execlastid
INSERT INTO users (name, dob) VALUES (?, ?);

-- name: GetUser :one
SELECT id, name, dob FROM users WHERE id = ? LIMIT 1;

-- name: ListUsers :many
SELECT id, name, dob FROM users;

-- name: UpdateUser :exec
UPDATE users SET name = ?, dob = ? WHERE id = ?;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = ?;