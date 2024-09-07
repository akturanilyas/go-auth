package types

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go-test/enums"
)

type Route struct {
	Method    string
	Path      enums.Endpoint
	Handler   func(c *fiber.Ctx) error
	Validator func(c *fiber.Ctx) error
}

type (
	ErrorResponse struct {
		Error       bool
		FailedField string
		Tag         string
		Value       interface{}
	}

	XValidator struct {
		Validator *validator.Validate
	}

	GlobalErrorHandlerResp struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}
)
