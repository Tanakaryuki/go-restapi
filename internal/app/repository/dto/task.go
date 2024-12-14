package dto

import (
	"database/sql"
)

type Task struct {
	ID                string         `json:"id"`
	Title             string         `json:"title"`
	Detail            string         `json:"description"`
	AdministratorUser sql.NullString `json:"administrator_user"`
	CreatedAt         string         `json:"created_at"`
	UpdatedAt         string         `json:"updated_at"`
}
