package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tedysaputro/book-catalog-with-go/src/author"
	"github.com/tedysaputro/book-catalog-with-go/src/hello"
	"github.com/tedysaputro/book-catalog-with-go/src/publisher"
)

// SetupRoutes configures all application routes
func SetupRoutes(app *fiber.App) {
	// Initialize services
	helloService := hello.NewHelloService()
	authorService := author.NewAuthorService()
	publisherService := publisher.NewPublisherService()

	// Initialize handlers
	helloHandler := hello.NewHelloHandler(helloService)
	authorHandler := author.NewAuthorHandler(authorService)
	publisherHandler := publisher.NewPublisherHandler(publisherService)

	// Register routes from each module
	helloHandler.RegisterRoutes(app)
	authorHandler.RegisterRoutes(app)
	publisherHandler.RegisterRoutes(app)
}
