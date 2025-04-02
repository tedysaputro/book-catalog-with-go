package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initialize database
	InitDB()

	// Create Fiber app
	app := fiber.New(fiber.Config{
		AppName: "Book Catalog API",
	})

	// Setup routes
	SetupRoutes(app)

	// Start server
	log.Fatal(app.Listen(":8080"))
}
