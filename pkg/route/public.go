package route

import (
	"fiber-boilerplate/app/controller"
	"fiber-boilerplate/app/services/user"
	"fiber-boilerplate/database"

	"github.com/gofiber/fiber/v2"
)

func PublicRoute(router fiber.Router, db *database.DB) {
	userRepo := user.CreateRepository(db)
	userService := user.CreateService(userRepo)
	userController := controller.NewUserController(userService)

	router.Post("/register", userController.RegisterUser)
	router.Post("/login", userController.Login)
}
