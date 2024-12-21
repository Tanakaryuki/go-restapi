package task

import (
	"encoding/json"
	"net/http"

	"github.com/Tanakaryuki/go-restapi/internal/app/handler/schema"
	"github.com/Tanakaryuki/go-restapi/internal/app/service/task"
	"github.com/Tanakaryuki/go-restapi/internal/domain/entity"
	"github.com/Tanakaryuki/go-restapi/pkg/errors"
	"github.com/Tanakaryuki/go-restapi/pkg/middleware"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	taskService task.TaskIService
	validate    *validator.Validate
}

func New(taskService task.TaskIService) *Handler {
	return &Handler{
		taskService: taskService,
		validate:    validator.New(),
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

func (h *Handler) CreateTask() func(http.ResponseWriter, *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {
		username := r.Context().Value(middleware.UsernameKey).(string)
		var req schema.CreateTaskRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			return err
		}
		if err := h.validate.Struct(req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return nil
		}
		t := &entity.Task{
			ID:                req.ID,
			Title:             req.Title,
			Detail:            req.Detail,
			AdministratorUser: username,
		}
		if err := h.taskService.CreateTask(r.Context(), t); err != nil {
			switch err.Error() {
			case errors.ErrIDInUse:
				http.Error(w, err.Error(), http.StatusBadRequest)
				return nil
			default:
				return err
			}
		}
		w.WriteHeader(http.StatusCreated)
		return nil
	}
}
