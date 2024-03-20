package middleware

import (
	"errors"
	"fiber-boilerplate/app/models"
	"fiber-boilerplate/pkg/config"
	"time"

	"github.com/gofiber/fiber/v2"
)

func SetCookies(c *fiber.Ctx, userData models.User) error {
	// Convert data to string if it's not already
	refreshToken, err := GenerateNewRefreshToken(userData)
	if err != nil {
		return err
	}
	c.Cookie(&fiber.Cookie{
		Name:    "refreshToken",
		Value:   refreshToken,
		Expires: time.Now().Add(time.Hour * time.Duration(config.AppConfig().JWTSecreteExpireHoursCount)),
	})
	return nil
}

func GetTokenCookie(c *fiber.Ctx) (string, error) {
	cookie := c.Cookies("refreshToken")
	if cookie == "" {
		return "", errors.New("token cookie not found")
	}
	return cookie, nil
}
