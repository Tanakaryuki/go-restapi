package task

import (
	"context"

	"github.com/Tanakaryuki/go-restapi/internal/domain/entity"
)

type IRepository interface {
	GetByID(ctx context.Context, id string) (*entity.Task, error)
}
