package route

import (
	"github.com/gofiber/fiber/v2"
)

var routes []Route

func SetRoutes(_route []Route) {
	routes = _route
}

func BuildRoutes(app *fiber.App) {
	for _, route := range routes {
		handlers := append(route.Middlewares, route.Handler)
		app.Add(route.Method, string(route.Path), handlers...)
	}
}
