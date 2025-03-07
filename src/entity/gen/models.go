// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package entitygen

import (
	"database/sql"
	"time"
)

type Todo struct {
	ID          int64          `db:"id" json:"id"`
	Title       string         `db:"title" json:"title"`
	Description string         `db:"description" json:"description"`
	CreatedAt   time.Time      `db:"created_at" json:"created_at"`
	CreatedBy   string         `db:"created_by" json:"created_by"`
	UpdatedAt   sql.NullTime   `db:"updated_at" json:"updated_at"`
	UpdatedBy   sql.NullString `db:"updated_by" json:"updated_by"`
	DeletedAt   sql.NullTime   `db:"deleted_at" json:"deleted_at"`
	DeletedBy   sql.NullString `db:"deleted_by" json:"deleted_by"`
	IsDeleted   int8           `db:"is_deleted" json:"is_deleted"`
}
