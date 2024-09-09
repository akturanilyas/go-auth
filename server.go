package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go-test/pkg/abstract"
	"go-test/pkg/auth"
	"go-test/pkg/database"
	"go-test/pkg/route"
	"go-test/pkg/user"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusBadRequest).JSON(abstract.GlobalErrorHandlerResp{
				Success: false,
				Message: err.Error(),
			})
		},
	})

	conn := database.Connect()

	if err := conn.AutoMigrate(&user.User{}); err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}

	println(conn)
	auth.SetDBConnection(conn)

	route.BuildRoutes(app)

	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}
}
