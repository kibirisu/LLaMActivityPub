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
