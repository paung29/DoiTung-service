package preharvest

import (
	"github.com/doitung/DoiTung-service/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func PreHarvestRoute(app *fiber.App, handler *preHarvestHandler) {
	preHarvestGroup := app.Group("preHarvest")

	preHarvestGroup.Post("/create", middleware.RequiredAuth, handler.CreateOrUpdatePreHarvestForm)
	preHarvestGroup.Get("/get-preHarvest-form", middleware.RequiredAuth, handler.GetPreHarvestFormDetails)
	preHarvestGroup.Get("/get-preHarvest-form-histories", middleware.RequiredAuth, handler.GetPreHarvestFormHistories)
	preHarvestGroup.Put("/update-preHarvest-form", middleware.RequiredAuth, handler.CreateOrUpdatePreHarvestForm)
}
