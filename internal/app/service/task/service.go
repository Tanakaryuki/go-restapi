package task

import (
	"context"
	"errors"

	"github.com/Tanakaryuki/go-restapi/internal/domain/entity"
	"github.com/Tanakaryuki/go-restapi/internal/domain/repository/task"
	pkgErrors "github.com/Tanakaryuki/go-restapi/pkg/errors"
)

type TaskIService interface {
	GetTasks(ctx context.Context, id string) (*entity.Task, error)
	CreateTask(ctx context.Context, task *entity.Task) error
}

type Service struct {
	taskRepository task.IRepository
}

func New(taskRepository task.IRepository) TaskIService {
	return &Service{
		taskRepository: taskRepository,
	}
}

func (s *Service) GetTasks(ctx context.Context, id string) (*entity.Task, error) {
	task, err := s.taskRepository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (s *Service) CreateTask(ctx context.Context, task *entity.Task) error {
	exists, err := s.taskRepository.ExistsByID(ctx, task.ID)
	if err != nil {
		return err
	}
	if exists {
		return errors.New(pkgErrors.ErrIDInUse)
	}
	if err = s.taskRepository.Create(ctx, task); err != nil {
		return err
	}
	return nil
}
