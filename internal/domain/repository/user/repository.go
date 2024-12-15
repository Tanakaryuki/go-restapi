package user

import (
	"context"

	"github.com/Tanakaryuki/go-restapi/internal/domain/entity"
)

type IRepository interface {
	CreateUser(ctx context.Context, user *entity.User) error
	ExistsByEmail(ctx context.Context, email string) (bool, error)
	ExistsByUsername(ctx context.Context, username string) (bool, error)
}
