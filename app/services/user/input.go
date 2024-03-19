package user

type RegisterInput struct {
	Email     string `json:"email" validate:"required,email"`
	Username  string `json:"username" validate:"required,lte=255"`
	Password  string `json:"password" validate:"required,lte=255"`
	FirstName string `json:"first_name" validate:"required,lte=255"`
	LastName  string `json:"last_name" validate:"required,lte=255"`
}

type LoginInput struct {
	Query    string `json:"query" validate:"required"`
	Password string `json:"password" validate:"required"`
}
