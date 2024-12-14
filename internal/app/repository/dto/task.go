package dto

import (
	"database/sql"
)

type Task struct {
	ID                string         `db:"id"`
	Title             string         `db:"title"`
	Detail            string         `db:"detail"`
	AdministratorUser sql.NullString `db:"administrator_user"`
	CreatedAt         string         `db:"created_at"`
	UpdatedAt         string         `db:"updated_at"`
}
