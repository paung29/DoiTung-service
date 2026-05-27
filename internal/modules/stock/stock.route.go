package stock

import (
	"github.com/doitung/DoiTung-service/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, handler *handler) {

	stock := app.Group("/stocks")

	stock.Post("/create-carry-over", middleware.RequiredAuth, middleware.RequireRoles("ADMIN"), handler.CreateCarryOver)
	stock.Post("/create-incoming", middleware.RequiredAuth, middleware.RequireRoles("ADMIN"), handler.CreateIncomingStock)
	stock.Post("/create-issued", middleware.RequiredAuth, middleware.RequireRoles("ADMIN"), handler.CreateIssuedStock)
}
