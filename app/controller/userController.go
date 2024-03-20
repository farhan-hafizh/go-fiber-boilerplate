package controller

import (
	"fiber-boilerplate/app/services/user"
	"fiber-boilerplate/pkg/helper"
	"fiber-boilerplate/pkg/middleware"

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
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	newUser, err := ctr.userService.RegisterUser(*input)

	if err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(newUser)
}

func (ctr *userController) Login(c *fiber.Ctx) error {
	input := &user.LoginInput{}

	if err := c.BodyParser(input); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
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
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}
	err = middleware.SetCookies(c, loggedUser)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": err.Error(),
		})
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
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": "Invalid refresh token",
		})
	}

	decryptedToken, err := helper.Decrypt(refreshToken)

	if err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": "Invalid refresh token",
		})
	}

	// Parse and validate the decrypted token
	user, err := helper.ValidateToken(c, decryptedToken, "refresh_token")
	if err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": "Invalid refresh token",
		})
	}

	token, _ := middleware.GenerateNewAccessToken(*user)

	response := map[string]interface{}{
		"token": token,
	}

	return middleware.SendResponse(c, fiber.StatusOK, response)
}
