package middleware

import (
	"fiber-boilerplate/app/models"
	"fiber-boilerplate/pkg/config"
	"fiber-boilerplate/pkg/helper"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
		parsedToken, err := jwt.Parse(string(decryptedToken), func(token *jwt.Token) (interface{}, error) {
			return []byte(config.AppConfig().JWTSecretKey), nil
		})
		if err != nil || !parsedToken.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		claims, ok := parsedToken.Claims.(jwt.MapClaims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		// Extract user claims
		userClaims := claims["user"].(map[string]interface{})

		// Convert credit balance to int32
		creditBalanceFloat, _ := userClaims["credit_balance"].(float64)
		userIDString := userClaims["id"].(string)
		userId, err := primitive.ObjectIDFromHex(userIDString)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		// Create User object
		user := models.User{
			ID:            userId,
			Email:         userClaims["email"].(string),
			Username:      userClaims["username"].(string),
			Password:      userClaims["password"].(string),
			Photo:         userClaims["photo"].(string),
			FirstName:     userClaims["first_name"].(string),
			LastName:      userClaims["last_name"].(string),
			CreditBalance: int32(creditBalanceFloat),
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
