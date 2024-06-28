// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: users.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createUser = `-- name: CreateUser :exec
INSERT INTO
  users ("login", "password", "user_role_id", firstname)
VALUES
  ($1, $2, 1,$1)
`

type CreateUserParams struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) error {
	_, err := q.db.Exec(ctx, createUser, arg.Login, arg.Password)
	return err
}

const getClaimsByLogin = `-- name: GetClaimsByLogin :one
SELECT
  users.id,
  user_roles.title
FROM
  users
  LEFT JOIN user_roles ON user_roles.id = users.user_role_id
WHERE
  users.login = $1
`

type GetClaimsByLoginRow struct {
	ID    int64       `json:"id"`
	Title pgtype.Text `json:"title"`
}

func (q *Queries) GetClaimsByLogin(ctx context.Context, login string) (GetClaimsByLoginRow, error) {
	row := q.db.QueryRow(ctx, getClaimsByLogin, login)
	var i GetClaimsByLoginRow
	err := row.Scan(&i.ID, &i.Title)
	return i, err
}

const getPasswordByLogin = `-- name: GetPasswordByLogin :one
SELECT
  PASSWORD
FROM
  users
WHERE
  login = $1
`

func (q *Queries) GetPasswordByLogin(ctx context.Context, login string) (string, error) {
	row := q.db.QueryRow(ctx, getPasswordByLogin, login)
	var password string
	err := row.Scan(&password)
	return password, err
}
