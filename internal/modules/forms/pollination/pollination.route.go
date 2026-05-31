package pollination

import (
	"github.com/doitung/DoiTung-service/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func PollinationRoute(app *fiber.App, handler *PollinationHandler) {
	pollination := app.Group("/pollinations")

	pollination.Post("/create", middleware.RequiredAuth, handler.CreateOrUpdatePollinationForm)
	pollination.Get("/get-pollination-form", middleware.RequiredAuth, handler.GetPollinationFormDetails)
	pollination.Get("/get-pollination-form-histories", middleware.RequiredAuth, handler.GetPollinationFormHistories)
	pollination.Put("/update-pollination-form", middleware.RequiredAuth, handler.CreateOrUpdatePollinationForm)
	pollination.Get("/get-pollination-forms-by-zone", middleware.RequiredAuth, middleware.RequireRoles("ADMIN"), handler.GetPollinationFormsByZoneId)
}
