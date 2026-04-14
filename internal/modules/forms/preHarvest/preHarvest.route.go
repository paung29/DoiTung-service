package preharvest

import (
	"github.com/doitung/DoiTung-service/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func PreHarvestRoute(app *fiber.App, handler *preHarvestHandler) {
	preHarvestGroup := app.Group("pre-harvest")

	preHarvestGroup.Post("/create", middleware.RequiredAuth, handler.CreateOrUpdatePreHarvestForm)
}
