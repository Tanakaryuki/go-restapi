package user

import (
	"context"
	"errors"

	"github.com/Tanakaryuki/go-restapi/internal/domain/entity"
	"github.com/Tanakaryuki/go-restapi/internal/domain/repository/user"
	"github.com/Tanakaryuki/go-restapi/pkg/auth"
	pkgErrors "github.com/Tanakaryuki/go-restapi/pkg/errors"
	"github.com/google/uuid"
)

type UserIService interface {
	CreateUser(ctx context.Context, user *entity.User) error
	Login(ctx context.Context, login *entity.User) (*entity.Token, error)
	GetUserByUsername(ctx context.Context, username string) (*entity.User, error)
}

type Service struct {
	userRepository user.IRepository
}

func New(userRepository user.IRepository) UserIService {
	return &Service{
		userRepository: userRepository,
	}
}

func (s *Service) CreateUser(ctx context.Context, user *entity.User) error {
	exists, err := s.userRepository.ExistsByEmail(ctx, user.Email)
	if err != nil {
		return err
	}
	if exists {
		return errors.New(pkgErrors.ErrEmailInUse)
	}
	exists, err = s.userRepository.ExistsByUsername(ctx, user.Username)
	if err != nil {
		return err
	}
	if exists {
		return errors.New(pkgErrors.ErrUsernameInUse)
	}

	user.UUID = uuid.New().String()
	hashedPassword, err := auth.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	err = s.userRepository.CreateUser(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Login(ctx context.Context, login *entity.User) (*entity.Token, error) {
	user, err := s.userRepository.GetByUsername(ctx, login.Username)
	if err != nil {
		return nil, err
	}
	if err = auth.VerifyPassword(user.Password, login.Password); err != nil {
		return nil, errors.New(pkgErrors.ErrInvalidPassword)
	}
	token, err := auth.CreateToken(user.Username)
	if err != nil {
		return nil, err
	}
	t := &entity.Token{
		Token: token,
	}
	return t, nil
}

func (s *Service) GetUserByUsername(ctx context.Context, username string) (*entity.User, error) {
	user, err := s.userRepository.GetByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	return user, nil
}
