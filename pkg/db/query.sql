-- name: GetUsers :many
SELECT * FROM users;

-- name: AddUser :exec
INSERT INTO users (
  username,
  password_hash,
  bio,
  followers_count,
  following_count,
  is_admin
) VALUES ($1, $2, $3, $4, $5, $6);

-- name: GetUser :one
SELECT * FROM users WHERE id = $1;

-- name: UpdateUser :exec
UPDATE users SET password_hash = $2, bio = $3, followers_count = $4, following_count = $5, is_admin = $6 WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;

-- name: GetPosts :many
SELECT * FROM posts;

-- name: AddPost :exec
INSERT INTO posts (user_id, content) VALUES ($1, $2);

-- name: GetPost :one
SELECT * FROM posts WHERE id = $1;

-- name: UpdatePost :exec
UPDATE posts SET content = $2, like_count = $3, share_count = $4, comment_count = $5 WHERE id = $1;

-- name: DeletePost :exec
DELETE FROM posts WHERE id = $1;
