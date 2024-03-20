package middleware

import (
	"fiber-boilerplate/app/models"

	"github.com/gofiber/fiber/v2"
)

func SendAuthorizedResponse(c *fiber.Ctx, statusCode int, data interface{}) error {
	user := c.Locals("user").(*models.User)
	accessToken, err := GenerateNewAccessToken(*user)
	if err != nil {
		return err
	}

	return c.Status(statusCode).JSON(fiber.Map{
		"data":  data,
		"token": accessToken,
	})
}

func SendResponse(c *fiber.Ctx, statusCode int, data interface{}) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"data": data,
	})
}
