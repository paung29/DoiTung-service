package harvestgrading

import (
	"github.com/doitung/DoiTung-service/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func HarvestGradingRoute(app *fiber.App, handler *HarvestGradingHandler) {
	harvestGradingGroup := app.Group("/harvest-grading")

	harvestGradingGroup.Post("/create", middleware.RequiredAuth, handler.CreateOrUpdateHarvestGradingForm)
	harvestGradingGroup.Get("/get-harvest-grading-form", middleware.RequiredAuth, handler.GetHarvestGradingFormDetails)
	harvestGradingGroup.Get("/get-harvest-grading-form-histories", middleware.RequiredAuth, handler.GetHarvestGradingFormHistories)
	harvestGradingGroup.Put("/update-harvest-grading-form", middleware.RequiredAuth, handler.CreateOrUpdateHarvestGradingForm)
	harvestGradingGroup.Get("/get-harvest-grading-forms-by-zone", middleware.RequiredAuth, middleware.RequireRoles("ADMIN"), handler.GetHarvestGradingFormsByZoneId)
}
