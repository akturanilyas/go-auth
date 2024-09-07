package builders

import (
	"github.com/gofiber/fiber/v2"
	"go-test/routes"
	"go-test/types"
)

func BuildRoutes(app *fiber.App) {
	for _, route := range routes.AuthRoutes {
		app.Add(route.Method, string(route.Path), route.Handler)
	}
}

func routeWrapper(handler func(c *fiber.Ctx) error, route types.Route) func(c *fiber.Ctx) error {
	if route.Validator != nil {
		return func(c *fiber.Ctx) error {
			if err := route.Validator(c); err != nil {
				return err
			}

			return handler(c)
		}
	}

	return func(c *fiber.Ctx) error {
		return handler(c)
	}
}
