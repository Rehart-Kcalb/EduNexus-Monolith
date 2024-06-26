// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: categories.sql

package db

import (
	"context"
)

const allCategories = `-- name: AllCategories :many
SELECT
  "name",
  "color"
FROM
  categories
`

type AllCategoriesRow struct {
	Name  string `json:"name"`
	Color int32  `json:"color"`
}

func (q *Queries) AllCategories(ctx context.Context) ([]AllCategoriesRow, error) {
	rows, err := q.db.Query(ctx, allCategories)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []AllCategoriesRow{}
	for rows.Next() {
		var i AllCategoriesRow
		if err := rows.Scan(&i.Name, &i.Color); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCategoryId = `-- name: GetCategoryId :one
select id from categories where name=$1 limit 1
`

func (q *Queries) GetCategoryId(ctx context.Context, name string) (int64, error) {
	row := q.db.QueryRow(ctx, getCategoryId, name)
	var id int64
	err := row.Scan(&id)
	return id, err
}
