package category

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// CategoryHandler handles HTTP requests for categories
type CategoryHandler struct {
	service CategoryService
}

// NewCategoryHandler creates a new instance of CategoryHandler
func NewCategoryHandler(service CategoryService) *CategoryHandler {
	return &CategoryHandler{service: service}
}

// CreateCategory handles POST /categories request
func (h *CategoryHandler) CreateCategory(c *fiber.Ctx) error {
	var request CategoryRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	category, err := h.service.CreateCategory(request)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(category)
}

// GetCategory handles GET /categories/:id request
func (h *CategoryHandler) GetCategory(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid category ID",
		})
	}

	category, err := h.service.GetCategory(uint(id))
	if err != nil {
		if err.Error() == "category not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(category)
}

// GetCategories handles GET /categories request
func (h *CategoryHandler) GetCategories(c *fiber.Ctx) error {
	page := uint(c.QueryInt("pages", 1))
	limit := uint(c.QueryInt("limit", 10))
	sortBy := c.Query("sortBy", "id")
	direction := c.Query("direction", "asc")
	categoryName := c.Query("categoryName", "")

	categories, err := h.service.GetCategories(page, limit, sortBy, direction, categoryName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(categories)
}

// UpdateCategory handles PUT /categories/:id request
func (h *CategoryHandler) UpdateCategory(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid category ID",
		})
	}

	var request CategoryRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	category, err := h.service.UpdateCategory(uint(id), request)
	if err != nil {
		if err.Error() == "category not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(category)
}

// DeleteCategory handles DELETE /categories/:id request
func (h *CategoryHandler) DeleteCategory(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid category ID",
		})
	}

	if err := h.service.DeleteCategory(uint(id)); err != nil {
		if err.Error() == "category not found" {
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

// RegisterRoutes registers the category routes
func (h *CategoryHandler) RegisterRoutes(app *fiber.App) {
	categories := app.Group("/categories")
	categories.Post("/", h.CreateCategory)
	categories.Get("/", h.GetCategories)
	categories.Get("/:id", h.GetCategory)
	categories.Put("/:id", h.UpdateCategory)
	categories.Delete("/:id", h.DeleteCategory)
}
