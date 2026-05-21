package warehouse

import (
	"github.com/doitung/DoiTung-service/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func WarehouseRoute(app *fiber.App, handler *WarehouseHandler) {
	warehouse := app.Group("/warehouses")

	warehouse.Post("/create", middleware.RequiredAuth, middleware.RequireRoles("ADMIN"), handler.CreateWarehouse)
	warehouse.Get("/get-all-warehouses", middleware.RequiredAuth, middleware.RequireRoles("ADMIN"), handler.GetWarehouses)
	warehouse.Get("/get-warehouse-by-id", middleware.RequiredAuth, middleware.RequireRoles("ADMIN"), handler.GetWarehouseById)
}
