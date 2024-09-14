package main

import (
	"auth-go/pkg/abstract"
	"auth-go/pkg/database"
	"auth-go/pkg/route"
	"auth-go/pkg/user"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
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

	database.SetDBConnection(conn)

	route.SetRoutes(Routes)
	route.BuildRoutes(app)

	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}
}
