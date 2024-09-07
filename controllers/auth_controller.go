package authController

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go-test/services/authservice"
	"go-test/types"
	"go-test/validations"
)

var Validate = validator.New()

func Login(c *fiber.Ctx) error {
	return c.SendString("Login")
}

func Register(c *fiber.Ctx) error {
	myValidator := &types.XValidator{
		Validator: Validate,
	}

	user := &validations.CreateUserValidation{
		Name:     c.Query("name"),
		Password: c.Query("password"),
	}

	// TODO validate user input better way
	if err := myValidator.Validator.Struct(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Validation failed",
			"errors":  err.Error(),
		})
	}

	response := authservice.Register(user)

	return c.Status(fiber.StatusCreated).JSON(response)
}
