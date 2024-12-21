package task

import (
	"context"

	"github.com/Tanakaryuki/go-restapi/internal/domain/entity"
)

type IRepository interface {
	GetByID(ctx context.Context, id string) (*entity.Task, error)
	ExistsByID(ctx context.Context, id string) (bool, error)
	Create(ctx context.Context, task *entity.Task) error
}
