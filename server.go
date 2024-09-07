package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go-test/builders"
	"go-test/models"
	"go-test/services/authservice"
	"go-test/services/databaseservice"
	"go-test/types"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusBadRequest).JSON(types.GlobalErrorHandlerResp{
				Success: false,
				Message: err.Error(),
			})
		},
	})

	conn := databaseservice.Connect()

	if err := conn.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}

	println(conn)
	authservice.SetDBConnection(conn)

	builders.BuildRoutes(app)

	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}
}
