package task

import (
	"context"

	"github.com/Tanakaryuki/go-restapi/internal/domain/entity"
)

type IRepository interface {
	GetByUsername(ctx context.Context, username string) ([]*entity.Task, error)
}
