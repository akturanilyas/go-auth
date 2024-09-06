package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go-test/builders"
	"go-test/models"
	databaseService "go-test/services"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	app := fiber.New()
	conn := databaseService.Connect()
	// Migrate the schema
	if err := conn.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}
	builders.BuildRoutes(app)

	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}
}
