package helper

import (
	"fiber-boilerplate/app/models"
	"fiber-boilerplate/pkg/config"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ValidateToken(c *fiber.Ctx, decryptedToken []byte, tokenType string) (*models.User, error) {
	secret := config.AppConfig().AccessTokenSecret
	if tokenType == "refresh_token" {
		secret = config.AppConfig().RefreshTokenSecret
	}
	parsedToken, err := parseJWT(decryptedToken, secret)
	if err != nil || !parsedToken.Valid {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}

	user, err := extractUserFromClaims(claims)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}

	return user, nil
}

func parseJWT(token []byte, secret string) (*jwt.Token, error) {
	return jwt.Parse(string(token), func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
}

func extractUserFromClaims(claims jwt.MapClaims) (*models.User, error) {
	userClaims, ok := claims["user"].(map[string]interface{})
	if !ok {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}

	creditBalanceFloat, _ := userClaims["credit_balance"].(float64)
	userIDString := userClaims["id"].(string)
	userID, err := primitive.ObjectIDFromHex(userIDString)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}

	user := models.User{
		ID:            userID,
		Email:         userClaims["email"].(string),
		Username:      userClaims["username"].(string),
		Password:      userClaims["password"].(string),
		Photo:         userClaims["photo"].(string),
		FirstName:     userClaims["first_name"].(string),
		LastName:      userClaims["last_name"].(string),
		CreditBalance: int32(creditBalanceFloat),
	}

	return &user, nil
}
