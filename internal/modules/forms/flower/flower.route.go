package flower

import (
	"github.com/doitung/DoiTung-service/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func FlowerRoutes(app *fiber.App, handler *FlowerHandler) {
	flower := app.Group("/flowers")

	flower.Post("/create", middleware.RequiredAuth, handler.CreateOrUpdateFlowerForm)
}
