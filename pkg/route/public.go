package route

import (
	"go-fiber-boilerplate/app/controller"
	"go-fiber-boilerplate/app/services/user"
	"go-fiber-boilerplate/database"

	"github.com/gofiber/fiber/v2"
)

func PublicRoute(router fiber.Router, db *database.DB) {
	userRepo := user.CreateRepository(db)
	userService := user.CreateService(userRepo)
	userController := controller.NewUserController(userService)

	router.Post("/register", userController.RegisterUser)
	router.Post("/login", userController.Login)
	router.Get("/refresh-token", userController.RefreshToken)
}
