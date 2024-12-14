package task

import (
	"encoding/json"
	"net/http"

	"github.com/Tanakaryuki/go-restapi/internal/app/handler/schema"
	"github.com/Tanakaryuki/go-restapi/internal/app/service/task"
)

type Handler struct {
	taskService task.TaskIService
}

func New(taskService task.TaskIService) *Handler {
	return &Handler{
		taskService: taskService,
	}
}

func (h *Handler) GetTask() func(http.ResponseWriter, *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {
		username := r.PathValue("username")
		tasks, err := h.taskService.GetTasks(r.Context(), username)
		if err != nil {
			return err
		}
		var schemaTasks []schema.Task
		for _, task := range tasks {
			schemaTask := schema.Task{
				ID:                task.ID,
				Title:             task.Title,
				Detail:            task.Detail,
				AdministratorUser: task.AdministratorUser,
				CreatedAt:         task.CreatedAt,
				UpdatedAt:         task.UpdatedAt,
			}
			schemaTasks = append(schemaTasks, schemaTask)
		}

		res := schema.GetTaskResponse{
			Task: schemaTasks,
		}
		if err := json.NewEncoder(w).Encode(res); err != nil {
			return err
		}
		return nil
	}
}
