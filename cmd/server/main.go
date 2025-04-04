package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/tedysaputro/book-catalog-with-go/internal/infrastructure/database"
	"github.com/tedysaputro/book-catalog-with-go/internal/infrastructure/http"
)

func main() {
	// Initialize database
	database.InitDB()

	// Create Fiber app
	app := fiber.New(fiber.Config{
		AppName: "Book Catalog API",
	})

	// Setup routes
	http.SetupRoutes(app)

	// Start server
	log.Fatal(app.Listen(":8080"))
}
