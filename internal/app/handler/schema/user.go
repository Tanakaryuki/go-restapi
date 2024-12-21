package schema

type CreateUserRequest struct {
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required"`
	Username    string `json:"username" validate:"required"`
	DisplayName string `json:"display_name" validate:"required"`
	IsAdmin     bool   `json:"is_admin"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type Token struct {
	Token     string `json:"token"`
	TokenType string `json:"token_type"`
}

type UserDetailResponse struct {
	UUID        string `json:"uuid"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	DisplayName string `json:"display_name"`
	IsAdmin     bool   `json:"is_admin"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	DeletedAt   string `json:"deleted_at"`
}
