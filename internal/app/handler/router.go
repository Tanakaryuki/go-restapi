package handler

import (
	"github.com/Tanakaryuki/go-restapi/internal/app/handler/task"
)

type Root struct {
	TaskHander *task.Handler
}

func New(taskHandler *task.Handler) *Root {
	return &Root{
		TaskHander: taskHandler,
	}
}
