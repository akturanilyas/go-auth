package route

import (
	"github.com/gofiber/fiber/v2"
)

func BuildRoutes(app *fiber.App) {
	// TODO for now
	for _, route := range Routes {
		app.Add(route.Method, string(route.Path), route.Handler)
	}
}
