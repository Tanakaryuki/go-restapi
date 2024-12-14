package task

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Tanakaryuki/go-restapi/internal/app/repository/dto"
	"github.com/Tanakaryuki/go-restapi/internal/domain/entity"
	"github.com/Tanakaryuki/go-restapi/internal/domain/repository/task"
	"github.com/jmoiron/sqlx"
)

type repository struct {
	conn *sqlx.DB
}

func New(conn *sqlx.DB) task.IRepository {
	return &repository{
		conn: conn,
	}
}

func (r *repository) GetByID(ctx context.Context, id string) (*entity.Task, error) {
	var t dto.Task
	query := `SELECT id, title, detail, administrator_user, created_at, updated_at FROM tasks WHERE id = ?`
	err := r.conn.GetContext(ctx, &t, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	task := &entity.Task{
		ID:                t.ID,
		Title:             t.Title,
		Detail:            t.Detail,
		AdministratorUser: t.AdministratorUser.String,
		CreatedAt:         t.CreatedAt,
		UpdatedAt:         t.UpdatedAt,
	}
	return task, nil
}
