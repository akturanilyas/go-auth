package builders

import (
	"github.com/gofiber/fiber/v2"
	"go-test/routes"
)

func BuildRoutes(app *fiber.App) {
	for _, route := range routes.AuthRoutes {
		app.Add(route.Method, route.Path, route.Handler)
	}
}
