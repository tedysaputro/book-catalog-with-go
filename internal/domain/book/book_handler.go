package book

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// BookHandler handles HTTP requests for books
type BookHandler struct {
	service BookService
}

// NewBookHandler creates a new instance of BookHandler
func NewBookHandler(service BookService) *BookHandler {
	return &BookHandler{service: service}
}

// CreateBook handles POST /books request
func (h *BookHandler) CreateBook(c *fiber.Ctx) error {
	var request BookRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	book, err := h.service.CreateBook(request)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(book)
}

// GetBook handles GET /books/:id request
func (h *BookHandler) GetBook(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid book ID",
		})
	}

	book, err := h.service.GetBook(uint(id))
	if err != nil {
		if err.Error() == "book not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(book)
}

// GetBooks handles GET /books request
func (h *BookHandler) GetBooks(c *fiber.Ctx) error {
	page := uint(c.QueryInt("pages", 1))
	limit := uint(c.QueryInt("limit", 10))
	sortBy := c.Query("sortBy", "id")
	direction := c.Query("direction", "asc")
	title := c.Query("title", "")

	books, err := h.service.GetBooks(page, limit, sortBy, direction, title)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(books)
}

// UpdateBook handles PUT /books/:id request
func (h *BookHandler) UpdateBook(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid book ID",
		})
	}

	var request BookRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	book, err := h.service.UpdateBook(uint(id), request)
	if err != nil {
		if err.Error() == "book not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(book)
}

// DeleteBook handles DELETE /books/:id request
func (h *BookHandler) DeleteBook(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid book ID",
		})
	}

	if err := h.service.DeleteBook(uint(id)); err != nil {
		if err.Error() == "book not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// RegisterRoutes registers the book routes
func (h *BookHandler) RegisterRoutes(app *fiber.App) {
	books := app.Group("/api/v1/books")
	books.Post("/", h.CreateBook)
	books.Get("/", h.GetBooks)
	books.Get("/:id", h.GetBook)
	books.Put("/:id", h.UpdateBook)
	books.Delete("/:id", h.DeleteBook)
}
