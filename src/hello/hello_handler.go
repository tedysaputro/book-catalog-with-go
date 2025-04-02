package hello

import "github.com/gofiber/fiber/v2"

// HelloResponseDTO is removed as per the instruction

// HelloHandler handles HTTP requests for hello operations
type HelloHandler struct {
	service HelloService
}

// NewHelloHandler creates a new instance of HelloHandler
func NewHelloHandler(service HelloService) *HelloHandler {
	return &HelloHandler{
		service: service,
	}
}

// GetHello handles GET /hello request
func (h *HelloHandler) GetHello(c *fiber.Ctx) error {
	response := h.service.GetHelloMessage()
	return c.JSON(response)
}

// RegisterRoutes registers all routes for hello module
func (h *HelloHandler) RegisterRoutes(app *fiber.App) {
	// Group routes under /api/v1
	api := app.Group("/api/v1")
	
	// Hello routes
	api.Get("/hello", h.GetHello)
}