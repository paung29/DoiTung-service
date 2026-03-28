package zone

import (
	"github.com/doitung/DoiTung-service/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, handler *ZoneHandler) {

	account := app.Group("/zones")

	account.Post("/create",middleware.RequiredAuth, middleware.RequireRoles("ADMIN"), handler.CreateZone)

}