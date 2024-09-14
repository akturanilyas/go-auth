package route

import (
	"auth-go/enums"
	"github.com/gofiber/fiber/v2"
)

type Route struct {
	Method      string
	Path        enums.Endpoint
	Handler     func(c *fiber.Ctx) error
	Validator   func(c *fiber.Ctx) error
	Middlewares []fiber.Handler
}
