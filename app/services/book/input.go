package book

type CreateBookInput struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
}
