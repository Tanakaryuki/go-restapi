package task

import (
	"context"

	"github.com/Tanakaryuki/go-restapi/internal/domain/entity"
	"github.com/Tanakaryuki/go-restapi/internal/domain/repository/task"
)

type TaskIService interface {
	GetTasks(ctx context.Context, username string) ([]*entity.Task, error)
}

type Service struct {
	taskRepository task.IRepository
}

func New(taskRepository task.IRepository) TaskIService {
	return &Service{
		taskRepository: taskRepository,
	}
}

func (s *Service) GetTasks(ctx context.Context, username string) ([]*entity.Task, error) {
	tasks, err := s.taskRepository.GetByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
