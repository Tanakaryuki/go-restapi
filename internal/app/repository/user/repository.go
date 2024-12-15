package user

import (
	"context"

	"github.com/Tanakaryuki/go-restapi/internal/app/repository/dto"
	"github.com/Tanakaryuki/go-restapi/internal/domain/entity"
	"github.com/Tanakaryuki/go-restapi/internal/domain/repository/user"
	"github.com/jmoiron/sqlx"
)

type repository struct {
	conn *sqlx.DB
}

func New(conn *sqlx.DB) user.IRepository {
	return &repository{
		conn: conn,
	}
}

func (r *repository) CreateUser(ctx context.Context, user *entity.User) error {
	u := dto.User{
		UUID:           user.UUID,
		Username:       user.Username,
		Email:          user.Email,
		HashedPassword: user.Password,
		DisplayName:    user.DisplayName,
		IsAdmin:        user.IsAdmin,
	}
	query := `INSERT INTO users (uuid, username, email, hashed_password, display_name, is_admin) VALUES (:uuid, :username, :email, :hashed_password, :display_name, :is_admin)`
	_, err := r.conn.NamedExecContext(ctx, query, u)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE email = ?)`
	err := r.conn.GetContext(ctx, &exists, query, email)
	return exists, err
}

func (r *repository) ExistsByUsername(ctx context.Context, username string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE username = ?)`
	err := r.conn.GetContext(ctx, &exists, query, username)
	return exists, err
}

func (r *repository) GetByUsername(ctx context.Context, username string) (*entity.User, error) {
	var u dto.User
	query := `SELECT uuid, username, email, hashed_password, display_name, is_admin, created_at, updated_at, deleted_at FROM users WHERE username = ?`
	err := r.conn.GetContext(ctx, &u, query, username)
	if err != nil {
		return nil, err
	}
	user := &entity.User{
		UUID:        u.UUID,
		Username:    u.Username,
		Email:       u.Email,
		Password:    u.HashedPassword,
		DisplayName: u.DisplayName,
		IsAdmin:     u.IsAdmin,
	}
	return user, nil
}
