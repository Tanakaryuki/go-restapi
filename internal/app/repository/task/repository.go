package task

import (
	"context"

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

func (r *repository) GetByUsername(ctx context.Context, username string) ([]*entity.Task, error) {
	query := `SELECT id, title, detail, administrator_user, created_at, update_at FROM tasks WHERE administrator_user = ?`
	rows, err := r.conn.QueryContext(ctx, query, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*entity.Task
	for rows.Next() {
		var t dto.Task
		if err := rows.Scan(&t.ID, &t.Title, &t.Detail, &t.AdministratorUser, &t.CreatedAt, &t.UpdatedAt); err != nil {
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
		tasks = append(tasks, task)
	}
	return tasks, nil
}
