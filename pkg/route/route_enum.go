package route

import (
	"github.com/gofiber/fiber/v2"
	"go-test/enums"
)

type Route struct {
	Method    string
	Path      enums.Endpoint
	Handler   func(c *fiber.Ctx) error
	Validator func(c *fiber.Ctx) error
}
