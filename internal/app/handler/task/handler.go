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
		id := r.PathValue("id")
		task, err := h.taskService.GetTasks(r.Context(), id)
		if err != nil {
			return err
		}
		if task == nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return nil
		}
		schemaTask := schema.Task{
			ID:                task.ID,
			Title:             task.Title,
			Detail:            task.Detail,
			AdministratorUser: task.AdministratorUser,
			CreatedAt:         task.CreatedAt,
			UpdatedAt:         task.UpdatedAt,
		}

		res := schema.GetTaskResponse{
			Task: schemaTask,
		}
		if err := json.NewEncoder(w).Encode(res); err != nil {
			return err
		}
		return nil
	}
}
