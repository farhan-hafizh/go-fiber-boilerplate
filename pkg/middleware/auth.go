package middleware

import (
	"go-fiber-boilerplate/app/models"
	"go-fiber-boilerplate/pkg/config"
	"go-fiber-boilerplate/pkg/helper"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// JWTProtected func for specify route group with JWT authentication.
// See: https://github.com/gofiber/contrib/jwt
func JWTProtected() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Missing bearer token",
			})
		}

		// Extract the token from the request header or query parameter
		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == "" {
			token = c.Query("token")
		}

		// Decrypt the token
		decryptedToken, err := helper.Decrypt(token)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		// Parse and validate the decrypted token using JWT library
		user, err := helper.ValidateToken(c, decryptedToken, "access_token")
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		// Save user to locals
		c.Locals("user", user)

		// Token is valid, proceed with the next middleware or route handler
		return c.Next()
	}
}

func GenerateNewAccessToken(user models.User) (string, error) {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["user"] = user
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(config.AppConfig().JWTSecreteExpireMinutesCount)).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(config.AppConfig().AccessTokenSecret))
	if err != nil {
		return "", err
	}

	encrypted, err := helper.Encrypt([]byte(t))

	if err != nil {
		return "", err
	}

	return encrypted, nil
}

func GenerateNewRefreshToken(user models.User) (string, error) {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["user"] = user
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(config.AppConfig().JWTSecreteExpireHoursCount)).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(config.AppConfig().RefreshTokenSecret))
	if err != nil {
		return "", err
	}

	encrypted, err := helper.Encrypt([]byte(t))

	if err != nil {
		return "", err
	}

	return encrypted, nil
}
