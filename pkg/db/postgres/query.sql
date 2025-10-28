-- name: GetUsers :many
SELECT * FROM users;

-- name: CreateUser :exec
INSERT INTO users (name, email) VALUES ($1, $2);
