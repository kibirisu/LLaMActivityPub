-- name: GetUsersQuery :many
SELECT * FROM users;

-- name: AddUserQuery :exec
INSERT INTO users (
  username,
  password_hash,
  bio,
  followers_count,
  following_count,
  is_admin
) VALUES ($1, $2, $3, $4, $5, $6);

-- name: GetUserQuery :one
SELECT * FROM users WHERE id = $1;

-- name: UpdateUserQuery :exec
UPDATE users SET password_hash = $2, bio = $3, followers_count = $4, following_count = $5, is_admin = $6 WHERE id = $1;

-- name: DeleteUserQuery :exec
DELETE FROM users WHERE id = $1;
