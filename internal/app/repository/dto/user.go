package dto

import (
	"database/sql"
)

type User struct {
	UUID           string         `db:"uuid"`
	Username       string         `db:"username"`
	Email          string         `db:"email"`
	HashedPassword string         `db:"hashed_password"`
	DisplayName    string         `db:"display_name"`
	IsAdmin        bool           `db:"is_admin"`
	CreatedAt      string         `db:"created_at"`
	UpdatedAt      string         `db:"updated_at"`
	DeletedAt      sql.NullString `db:"deleted_at"`
}
