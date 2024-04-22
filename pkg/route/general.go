package route

import (
	"fiber-boilerplate/database"

	"github.com/gofiber/fiber/v2"
)

func GeneralRoute(a *fiber.App) {
	a.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"msg":    "Welcome to Fiber Go API!",
			"status": "/h34l7h",
		})
	})

	a.Get("/h34l7h", func(c *fiber.Ctx) error {
		err := database.PingDB()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"msg": err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"msg":       "Health Check",
			"db_online": true,
		})
	})
}

func NotFoundRoute(a *fiber.App) {
	a.Use(
		func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"msg": "sorry, endpoint is not found",
			})
		},
	)
}
