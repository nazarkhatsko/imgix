-- name: GetConfig :one
SELECT * FROM configs WHERE id = $1 LIMIT 1;

-- name: ListConfigs :many
SELECT * FROM configs ORDER BY id;
