package schema

type Task struct {
	ID                string `json:"id"`
	Title             string `json:"title"`
	Detail            string `json:"description"`
	AdministratorUser string `json:"administrator_user"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
}

type GetTaskResponse struct {
	Task []Task `json:"task"`
}