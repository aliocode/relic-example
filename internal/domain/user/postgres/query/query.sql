-- name: CreateUser :exec
INSERT INTO userdata (id, name, email) VALUES ($1, $2, $3);

-- name: FetchByEmail :one
SELECT * FROM userdata WHERE email = $1;