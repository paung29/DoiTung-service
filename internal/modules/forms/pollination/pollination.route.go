package pollination

import (
	"github.com/doitung/DoiTung-service/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func PollinationRoute(app *fiber.App, handler *PollinationHandler) {
	pollination := app.Group("/pollinations")

	pollination.Post("/create", middleware.RequiredAuth, handler.CreateOrUpdatePollinationForm)
	pollination.Get("/get-pollination-form", middleware.RequiredAuth, handler.GetPollinationFormDetails)
}
