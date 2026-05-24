package customer

import (
	"github.com/doitung/DoiTung-service/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func CustomerRoute(app *fiber.App, handler *CustomerHandler) {
	customer := app.Group("/customers")

	customer.Post("/create", middleware.RequiredAuth, middleware.RequireRoles("ADMIN"), handler.CreateCustomer)
}
