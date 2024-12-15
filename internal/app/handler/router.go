package handler

import (
	"github.com/Tanakaryuki/go-restapi/internal/app/handler/task"
	"github.com/Tanakaryuki/go-restapi/internal/app/handler/user"
)

type Root struct {
	TaskHander *task.Handler
	UserHander *user.Handler
}

func New(taskHandler *task.Handler, userHandler *user.Handler) *Root {
	return &Root{
		TaskHander: taskHandler,
		UserHander: userHandler,
	}
}
