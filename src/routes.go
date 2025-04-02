package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tedysaputro/book-catalog-with-go/src/author"
	"github.com/tedysaputro/book-catalog-with-go/src/hello"
)

// SetupRoutes configures all application routes
func SetupRoutes(app *fiber.App) {
	// Initialize services
	helloService := hello.NewHelloService()
	authorService := author.NewAuthorService()

	// Initialize handlers
	helloHandler := hello.NewHelloHandler(helloService)
	authorHandler := author.NewAuthorHandler(authorService)

	// Register routes from each module
	helloHandler.RegisterRoutes(app)
	authorHandler.RegisterRoutes(app)
}
