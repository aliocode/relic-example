// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: query.sql

package sqlgen

import (
	"context"
)

const createUser = `-- name: CreateUser :exec
INSERT INTO userdata (id, name, email) VALUES ($1, $2, $3)
`

type CreateUserParams struct {
	ID    string
	Name  string
	Email string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) error {
	_, err := q.db.Exec(ctx, createUser, arg.ID, arg.Name, arg.Email)
	return err
}

const fetchByEmail = `-- name: FetchByEmail :one
SELECT id, name, email, created_at, updated_at FROM userdata WHERE email = $1
`

func (q *Queries) FetchByEmail(ctx context.Context, email string) (Userdatum, error) {
	row := q.db.QueryRow(ctx, fetchByEmail, email)
	var i Userdatum
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}