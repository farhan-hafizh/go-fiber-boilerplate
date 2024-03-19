package middleware

import (
	"fiber-boilerplate/app/models"
	"fiber-boilerplate/pkg/config"
	"fiber-boilerplate/pkg/helper"
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

		if !strings.Contains(authHeader, "Bearer ") {
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

		// Now, you can parse and validate the decrypted token using JWT library
		claims := jwt.MapClaims{}
		_, err = jwt.ParseWithClaims(string(decryptedToken), claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.AppConfig().JWTSecretKey), nil
		})

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		// save user from token to locals
		c.Locals("user", claims)

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
	t, err := token.SignedString([]byte(config.AppConfig().JWTSecretKey))
	if err != nil {
		return "", err
	}

	encrypted, err := helper.Encrypt([]byte(t))

	if err != nil {
		return "", err
	}

	return encrypted, nil
}
