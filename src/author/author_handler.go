package author

import "github.com/gofiber/fiber/v2"

// AuthorHandler handles HTTP requests for author operations
type AuthorHandler struct {
	service AuthorService
}

// NewAuthorHandler creates a new instance of AuthorHandler
func NewAuthorHandler(service AuthorService) *AuthorHandler {
	return &AuthorHandler{
		service: service,
	}
}

// CreateAuthor handles POST /authors request
func (h *AuthorHandler) CreateAuthor(c *fiber.Ctx) error {
	var request AuthorCreateRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	dto, err := h.service.createAuthor(request)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(dto)
}

// GetAuthor handles GET /authors/:id request
func (h *AuthorHandler) GetAuthor(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid author ID",
		})
	}

	author, err := h.service.GetAuthor(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(author)
}

// GetAuthors handles GET /authors request
func (h *AuthorHandler) GetAuthors(c *fiber.Ctx) error {
	page := uint(c.QueryInt("p", 1))
	limit := uint(c.QueryInt("limit", 10))
	sortBy := c.Query("sortBy", "id")
	direction := c.Query("direction", "asc")
	authorName := c.Query("authorName", "")

	authors, err := h.service.GetAuthors(page, limit, sortBy, direction, authorName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(authors)
}

// RegisterRoutes registers all routes for author module
func (h *AuthorHandler) RegisterRoutes(app *fiber.App) {
	// Group routes under /app/v1
	v1 := app.Group("/api/v1")

	// Author routes
	authors := v1.Group("/authors")
	authors.Post("/", h.CreateAuthor)
	authors.Get("/", h.GetAuthors)
	authors.Get("/:id", h.GetAuthor)
}
