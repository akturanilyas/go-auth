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
		Value:    token,
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

	newUser, err := RegisterService(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not register user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(newUser)
}

func Logout(c *fiber.Ctx) error {
	// Implement token invalidation logic here if needed
	return c.SendStatus(fiber.StatusOK)
}

func RefreshToken(c *fiber.Ctx) error {
	// Implement token refresh logic here
	return c.SendStatus(fiber.StatusOK)
}

func ForgotPassword(c *fiber.Ctx) error {
	// Implement password reset logic here
	return c.SendStatus(fiber.StatusOK)
}
