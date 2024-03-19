package controller

import (
	"fiber-boilerplate/app/models"
	"fiber-boilerplate/app/services/book"

	"github.com/gofiber/fiber/v2"
)

type BookController interface {
	CreateBook(c *fiber.Ctx) error
}

type bookController struct {
	bookService book.Service
}

func NewBookController(bookService book.Service) BookController {
	return &bookController{bookService}
}

func (ctr *bookController) CreateBook(c *fiber.Ctx) error {

	// Create new Book struct
	input := &book.CreateBookInput{}

	if err := c.BodyParser(input); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	user := c.Locals("user").(models.User)

	newBook, err := ctr.bookService.CreateBook(*input, user)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"book": newBook,
	})
}
