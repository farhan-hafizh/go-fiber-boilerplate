package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func FiberMiddleware(app *fiber.App) {
	app.Use(
		// Add CORS to each route.
		cors.New(),
		// Add simple logger.
		logger.New(),
	)
}
