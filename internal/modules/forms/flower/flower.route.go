package flower

import (
	"github.com/doitung/DoiTung-service/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func FlowerRoutes(app *fiber.App, handler *FlowerHandler) {
	flower := app.Group("/flowers")

	flower.Post("/create", middleware.RequiredAuth, handler.CreateOrUpdateFlowerForm)
	flower.Get("/get-flower-form", middleware.RequiredAuth, handler.GetFlowerFormDetails)
	flower.Get("/get-flower-form-histories", middleware.RequiredAuth, handler.GetFlowerFormHistories)
	flower.Put("/update-flower-form", middleware.RequiredAuth, handler.CreateOrUpdateFlowerForm)
}
