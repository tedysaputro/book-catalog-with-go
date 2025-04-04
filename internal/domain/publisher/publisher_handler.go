package publisher

import (
	"github.com/gofiber/fiber/v2"
)

// PublisherHandler handles HTTP requests for publisher operations
type PublisherHandler struct {
	service PublisherService
}

// NewPublisherHandler creates a new instance of PublisherHandler
func NewPublisherHandler(service PublisherService) *PublisherHandler {
	return &PublisherHandler{service: service}
}

// CreatePublisher handles POST /publishers request
func (h *PublisherHandler) CreatePublisher(c *fiber.Ctx) error {
	var request PublisherRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if request.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Name is required",
		})
	}

	dto, err := h.service.createPublisher(request)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(dto)
}

// UpdatePublisher handles PUT /publishers/:id request
func (h *PublisherHandler) UpdatePublisher(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid publisher ID",
		})
	}

	var request PublisherRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if request.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Name is required",
		})
	}

	dto, err := h.service.UpdatePublisher(uint(id), request)
	if err != nil {
		if err.Error() == "publisher not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(dto)
}

// GetPublisher handles GET /publishers/:id request
func (h *PublisherHandler) GetPublisher(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid publisher ID",
		})
	}

	dto, err := h.service.GetPublisher(uint(id))
	if err != nil {
		if err.Error() == "publisher not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(dto)
}

// GetPublishers handles GET /publishers request
func (h *PublisherHandler) GetPublishers(c *fiber.Ctx) error {
	page := uint(c.QueryInt("pages", 1))
	limit := uint(c.QueryInt("limit", 10))
	sortBy := c.Query("sortBy", "id")
	direction := c.Query("direction", "asc")
	publisherName := c.Query("publisherName", "")

	publishers, err := h.service.GetPublishers(page, limit, sortBy, direction, publisherName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(publishers)
}

func (h *PublisherHandler) DeletePublisher(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid publisher ID",
		})
	}

	err = h.service.DeletePublisher(uint(id))
	if err != nil {
		if err.Error() == "publisher not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Publisher deleted successfully",
	})
}

// RegisterRoutes registers all routes for publisher module
func (h *PublisherHandler) RegisterRoutes(app *fiber.App) {
	// Group routes under /app/v1
	v1 := app.Group("/api/v1")

	// Publisher routes
	publishers := v1.Group("/publishers")
	publishers.Post("/", h.CreatePublisher)
	publishers.Get("/", h.GetPublishers)
	publishers.Get("/:id", h.GetPublisher)
	publishers.Put("/:id", h.UpdatePublisher)
	publishers.Delete("/:id", h.DeletePublisher)
}
