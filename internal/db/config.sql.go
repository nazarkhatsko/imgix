// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: config.sql

package db

import (
	"context"
)

const getConfig = `-- name: GetConfig :one
SELECT id, value FROM configs WHERE id = $1 LIMIT 1
`

func (q *Queries) GetConfig(ctx context.Context, id int64) (Config, error) {
	row := q.db.QueryRow(ctx, getConfig, id)
	var i Config
	err := row.Scan(&i.ID, &i.Value)
	return i, err
}

const listConfigs = `-- name: ListConfigs :many
SELECT id, value FROM configs ORDER BY id
`

func (q *Queries) ListConfigs(ctx context.Context) ([]Config, error) {
	rows, err := q.db.Query(ctx, listConfigs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Config
	for rows.Next() {
		var i Config
		if err := rows.Scan(&i.ID, &i.Value); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
