-- name: AddUser :one
INSERT INTO users (email, password, role) 
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE email = $1;

-- name: AddPost :one
INSERT INTO posts (title, content, category, tags) 
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: UpdatePost :one
UPDATE posts 
SET title = $2, content = $3, category = $4, tags = $5 
WHERE id = $1
RETURNING *;

-- name: DeletePost :one
DELETE FROM posts 
WHERE id = $1
RETURNING *;

-- name: ListPosts :many
SELECT * FROM posts;

-- name: GetPost :one
SELECT * FROM posts 
WHERE id = $1;

-- name: FilterPosts :many
SELECT * FROM posts 
WHERE title ILIKE '%' || $1 || '%' 
   OR content ILIKE '%' || $1 || '%' 
   OR category ILIKE '%' || $1 || '%';


