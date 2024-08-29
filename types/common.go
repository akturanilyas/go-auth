package types

import "github.com/gofiber/fiber/v2"

type Route struct {
	Method  string
	Path    string
	Handler func(c *fiber.Ctx) error
}
