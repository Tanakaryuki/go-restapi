package user

import (
	"encoding/json"
	"net/http"

	"github.com/Tanakaryuki/go-restapi/internal/app/handler/schema"
	"github.com/Tanakaryuki/go-restapi/internal/app/service/user"
	"github.com/Tanakaryuki/go-restapi/internal/domain/entity"
	"github.com/Tanakaryuki/go-restapi/pkg/errors"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	userService user.UserIService
	validate    *validator.Validate
}

func New(userService user.UserIService) *Handler {
	return &Handler{
		userService: userService,
		validate:    validator.New(),
	}
}

func (h *Handler) CreateUser() func(http.ResponseWriter, *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {
		req := schema.CreateUserRequest{}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return nil
		}
		if err := h.validate.Struct(req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return nil
		}

		u := &entity.User{
			Email:       req.Email,
			Password:    req.Password,
			Username:    req.Username,
			DisplayName: req.DisplayName,
			IsAdmin:     req.IsAdmin,
		}
		if err := h.userService.CreateUser(r.Context(), u); err != nil {
			switch err.Error() {
			case errors.ErrEmailInUse, errors.ErrUsernameInUse:
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

func (h *Handler) Login() func(http.ResponseWriter, *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {
		req := schema.LoginRequest{}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return nil
		}
		if err := h.validate.Struct(req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return nil
		}
		u := &entity.User{
			Username: req.Username,
			Password: req.Password,
		}

		token, err := h.userService.Login(r.Context(), u)
		if err != nil {
			switch err.Error() {
			case errors.ErrInvalidPassword:
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return nil
			default:
				return err
			}
		}
		res := schema.Token{
			Token:     token.Token,
			TokenType: "Bearer",
		}
		if err := json.NewEncoder(w).Encode(res); err != nil {
			return err
		}
		return nil
	}
}
