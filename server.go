package main

import (
	"github.com/gofiber/fiber/v2"
	"go-test/builders"
	"log"
)

func main() {
	app := fiber.New()

	builders.BuildRoutes(app)

	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}
}
