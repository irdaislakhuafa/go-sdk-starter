// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: queries.sql

package entitygen

import (
	"context"
	"database/sql"
	"time"
)

const countTodo = `-- name: CountTodo :one
SELECT
 COUNT(` + "`" + `id` + "`" + `) AS total
FROM
 ` + "`" + `todo` + "`" + `
WHERE
 ` + "`" + `is_deleted` + "`" + ` = ?
 AND (
  ` + "`" + `title` + "`" + ` LIKE CONCAT('%', ? , '%') OR ` + "`" + `description` + "`" + ` LIKE CONCAT('%', ?, '%')
 )
`

type CountTodoParams struct {
	IsDeleted int8        `db:"is_deleted" json:"is_deleted"`
	CONCAT    interface{} `db:"CONCAT" json:"CONCAT"`
	CONCAT_2  interface{} `db:"CONCAT_2" json:"CONCAT_2"`
}

func (q *Queries) CountTodo(ctx context.Context, arg CountTodoParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, countTodo, arg.IsDeleted, arg.CONCAT, arg.CONCAT_2)
	var total int64
	err := row.Scan(&total)
	return total, err
}

const createTodo = `-- name: CreateTodo :execresult
INSERT INTO ` + "`" + `todo` + "`" + ` (` + "`" + `title` + "`" + `, ` + "`" + `description` + "`" + `, ` + "`" + `created_at` + "`" + `, ` + "`" + `created_by` + "`" + `)
VALUES (?, ?, ?, ?)
`

type CreateTodoParams struct {
	Title       string    `db:"title" json:"title"`
	Description string    `db:"description" json:"description"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	CreatedBy   string    `db:"created_by" json:"created_by"`
}

func (q *Queries) CreateTodo(ctx context.Context, arg CreateTodoParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createTodo,
		arg.Title,
		arg.Description,
		arg.CreatedAt,
		arg.CreatedBy,
	)
}

const getTodo = `-- name: GetTodo :one
SELECT id, title, description, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by, is_deleted FROM ` + "`" + `todo` + "`" + ` WHERE ` + "`" + `id` + "`" + ` = ?
`

func (q *Queries) GetTodo(ctx context.Context, id int64) (Todo, error) {
	row := q.db.QueryRowContext(ctx, getTodo, id)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.UpdatedAt,
		&i.UpdatedBy,
		&i.DeletedAt,
		&i.DeletedBy,
		&i.IsDeleted,
	)
	return i, err
}

const listTodo = `-- name: ListTodo :many
SELECT id, title, description, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by, is_deleted
FROM ` + "`" + `todo` + "`" + `
WHERE
 ` + "`" + `is_deleted` + "`" + ` = ?
 AND (
  ` + "`" + `title` + "`" + ` LIKE CONCAT('%', ? , '%') OR ` + "`" + `description` + "`" + ` LIKE CONCAT('%', ?, '%')
 )
ORDER BY (
  ` + "`" + `title` + "`" + ` LIKE CONCAT('%', ? , '%') OR ` + "`" + `description` + "`" + ` LIKE CONCAT('%', ?, '%')
) DESC
LIMIT ?
OFFSET ?
`

type ListTodoParams struct {
	IsDeleted int8        `db:"is_deleted" json:"is_deleted"`
	CONCAT    interface{} `db:"CONCAT" json:"CONCAT"`
	CONCAT_2  interface{} `db:"CONCAT_2" json:"CONCAT_2"`
	CONCAT_3  interface{} `db:"CONCAT_3" json:"CONCAT_3"`
	CONCAT_4  interface{} `db:"CONCAT_4" json:"CONCAT_4"`
	Limit     int32       `db:"limit" json:"limit"`
	Offset    int32       `db:"offset" json:"offset"`
}

func (q *Queries) ListTodo(ctx context.Context, arg ListTodoParams) ([]Todo, error) {
	rows, err := q.db.QueryContext(ctx, listTodo,
		arg.IsDeleted,
		arg.CONCAT,
		arg.CONCAT_2,
		arg.CONCAT_3,
		arg.CONCAT_4,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Todo
	for rows.Next() {
		var i Todo
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.CreatedAt,
			&i.CreatedBy,
			&i.UpdatedAt,
			&i.UpdatedBy,
			&i.DeletedAt,
			&i.DeletedBy,
			&i.IsDeleted,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateTodo = `-- name: UpdateTodo :execresult
UPDATE ` + "`" + `todo` + "`" + ` SET
 ` + "`" + `title` + "`" + ` = ?,
 ` + "`" + `description` + "`" + ` = ?,
 ` + "`" + `updated_at` + "`" + ` = ?,
 ` + "`" + `updated_by` + "`" + ` = ?,
 ` + "`" + `deleted_at` + "`" + ` = ?,
 ` + "`" + `deleted_by` + "`" + ` = ?,
 ` + "`" + `is_deleted` + "`" + ` = ?
WHERE ` + "`" + `id` + "`" + ` = ?
`

type UpdateTodoParams struct {
	Title       string         `db:"title" json:"title"`
	Description string         `db:"description" json:"description"`
	UpdatedAt   sql.NullTime   `db:"updated_at" json:"updated_at"`
	UpdatedBy   sql.NullString `db:"updated_by" json:"updated_by"`
	DeletedAt   sql.NullTime   `db:"deleted_at" json:"deleted_at"`
	DeletedBy   sql.NullString `db:"deleted_by" json:"deleted_by"`
	IsDeleted   int8           `db:"is_deleted" json:"is_deleted"`
	ID          int64          `db:"id" json:"id"`
}

func (q *Queries) UpdateTodo(ctx context.Context, arg UpdateTodoParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateTodo,
		arg.Title,
		arg.Description,
		arg.UpdatedAt,
		arg.UpdatedBy,
		arg.DeletedAt,
		arg.DeletedBy,
		arg.IsDeleted,
		arg.ID,
	)
}
