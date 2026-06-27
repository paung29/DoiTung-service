package pole

import (
	"github.com/doitung/DoiTung-service/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func PoleRoute(app *fiber.App, handler *PoleHandler) {
	pole := app.Group("/poles")

	pole.Get("/get-by-zone", middleware.RequiredAuth, handler.GetPoleByZone)
	pole.Get("/get-pole-filter", middleware.RequiredAuth, handler.GetPoleFilter)
}
