package route

import (
	"go-fiber-boilerplate/app/controller"
	"go-fiber-boilerplate/app/services/book"
	"go-fiber-boilerplate/database"
	"go-fiber-boilerplate/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func PrivateRoute(router fiber.Router, db *database.DB) {
	bookRepo := book.CreateRepository(db)
	bookService := book.CreateService(bookRepo)
	bookController := controller.NewBookController(bookService)

	booksRoute := router.Group("/books", middleware.JWTProtected())

	booksRoute.Post("/create", bookController.CreateBook)

}
