package auth

import (
	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	loginData := new(LoginValidation)
	c.BodyParser(loginData)

	token, err := LoginService(loginData)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    "Bearer " + token,
		HTTPOnly: true,
		Secure:   true,
	})

	return c.JSON(fiber.Map{
		"token": token,
	})
}

func Register(c *fiber.Ctx) error {
	user := new(CreateUserValidation)

	c.BodyParser(user)

	newUser, _ := RegisterService(user)

	if newUser == nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "User already exists",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(newUser)
}

func Logout(c *fiber.Ctx) error {
	c.ClearCookie("jwt")

	return c.SendStatus(fiber.StatusOK)
}

func RefreshToken(c *fiber.Ctx) error {
	// Implement token refresh logic here
	return c.SendStatus(fiber.StatusOK)
}
