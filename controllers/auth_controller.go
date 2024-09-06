package authController

import "github.com/gofiber/fiber/v2"

func Login(c *fiber.Ctx) error {
	return c.SendString("Login")
}

func Register(c *fiber.Ctx) error {
	response := fiber.Map{
		"status":  "success",
		"message": "User registered successfully",
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}
