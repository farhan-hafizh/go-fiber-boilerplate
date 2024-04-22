package controller

import (
	"go-fiber-boilerplate/app/services/user"
	"go-fiber-boilerplate/pkg/helper"
	"go-fiber-boilerplate/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

type UserController interface {
	RegisterUser(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	RefreshToken(c *fiber.Ctx) error
}

type userController struct {
	userService user.Service
}

func NewUserController(userService user.Service) UserController {
	return &userController{userService}
}

func (ctr *userController) RegisterUser(c *fiber.Ctx) error {
	input := &user.RegisterInput{}

	if err := c.BodyParser(input); err != nil {
		return middleware.SendResponse(c, fiber.StatusBadRequest, err.Error())
	}

	newUser, err := ctr.userService.RegisterUser(*input)

	if err != nil {
		return middleware.SendResponse(c, fiber.StatusBadRequest, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(newUser)
}

func (ctr *userController) Login(c *fiber.Ctx) error {
	input := &user.LoginInput{}

	if err := c.BodyParser(input); err != nil {
		// Return status 400 and error message.
		return middleware.SendResponse(c, fiber.StatusBadRequest, err.Error())
	}

	loggedUser, err := ctr.userService.Login(*input)

	userDto := user.UserDto{}
	userDto.Username = loggedUser.Username
	userDto.Email = loggedUser.Email
	userDto.Photo = loggedUser.Photo
	userDto.FirstName = loggedUser.FirstName
	userDto.LastName = loggedUser.LastName
	userDto.CreditBalance = loggedUser.CreditBalance

	if err != nil {
		return middleware.SendResponse(c, fiber.StatusBadRequest, err.Error())

	}

	err = middleware.SetCookies(c, loggedUser)
	if err != nil {
		return middleware.SendResponse(c, fiber.StatusBadRequest, err.Error())

	}

	token, _ := middleware.GenerateNewAccessToken(loggedUser)
	response := map[string]interface{}{
		"user":  userDto,
		"token": token,
	}
	return middleware.SendResponse(c, fiber.StatusOK, response)

}

func (ctr *userController) RefreshToken(c *fiber.Ctx) error {
	refreshToken, err := middleware.GetTokenCookie(c)
	if err != nil {
		// Return status 400 and error message.
		return middleware.SendResponse(c, fiber.StatusBadRequest, "Invalid refresh token")

	}

	decryptedToken, err := helper.Decrypt(refreshToken)

	if err != nil {
		return middleware.SendResponse(c, fiber.StatusBadRequest, "Invalid refresh token")
	}

	// Parse and validate the decrypted token
	user, err := helper.ValidateToken(c, decryptedToken, "refresh_token")
	if err != nil {
		return middleware.SendResponse(c, fiber.StatusBadRequest, "Invalid refresh token")
	}

	token, _ := middleware.GenerateNewAccessToken(*user)

	response := map[string]interface{}{
		"token": token,
	}

	return middleware.SendResponse(c, fiber.StatusOK, response)
}
