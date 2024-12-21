package entity

type User struct {
	UUID        string
	Username    string
	Email       string
	DisplayName string
	Password    string
	IsAdmin     bool
	UpdatedAt   string
	CreatedAt   string
	DeletedAt   string
}

type Token struct {
	Token     string
	TokenType string
}
