package warehouse

import (
	"github.com/doitung/DoiTung-service/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func WarehouseRoute(app *fiber.App, handler *WarehouseHandler) {
	warehouse := app.Group("/warehouses")

	warehouse.Post("/create", middleware.RequiredAuth, middleware.RequireRoles("ADMIN"), handler.CreateWarehouse)
}
