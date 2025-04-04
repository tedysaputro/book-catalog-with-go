package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tedysaputro/book-catalog-with-go/internal/domain/author"
	"github.com/tedysaputro/book-catalog-with-go/internal/domain/book"
	"github.com/tedysaputro/book-catalog-with-go/internal/domain/category"
	"github.com/tedysaputro/book-catalog-with-go/internal/domain/hello"
	"github.com/tedysaputro/book-catalog-with-go/internal/domain/publisher"
)

// SetupRoutes configures all application routes
func SetupRoutes(app *fiber.App) {
	// Initialize services
	helloService := hello.NewHelloService()
	authorService := author.NewAuthorService()
	publisherService := publisher.NewPublisherService()
	categoryService := category.NewCategoryService()
	bookService := book.NewBookService()

	// Initialize handlers
	helloHandler := hello.NewHelloHandler(helloService)
	authorHandler := author.NewAuthorHandler(authorService)
	publisherHandler := publisher.NewPublisherHandler(publisherService)
	categoryHandler := category.NewCategoryHandler(categoryService)
	bookHandler := book.NewBookHandler(bookService)

	// Register routes from each module
	helloHandler.RegisterRoutes(app)
	authorHandler.RegisterRoutes(app)
	publisherHandler.RegisterRoutes(app)
	categoryHandler.RegisterRoutes(app)
	bookHandler.RegisterRoutes(app)
}
