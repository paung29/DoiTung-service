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
	// stock.Put("/update", middleware.RequiredAuth, middleware.RequireRoles("ADMIN"), handler.UpdateStockMovement)
	stock.Delete("/delete", middleware.RequiredAuth, middleware.RequireRoles("ADMIN"), handler.DeleteStockMovement)
	stock.Get("/get-all-by-year", middleware.RequiredAuth, middleware.RequireRoles("ADMIN"), handler.GetStockMovementListsByYear)
}
