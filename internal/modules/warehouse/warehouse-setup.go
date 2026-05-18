package warehouse

import (
	"github.com/doitung/DoiTung-service/internal/modules/year"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Setup(app *fiber.App, db *gorm.DB) {
	yearRepo := year.NewYearRepository(db)
	warehouseRepo := NewWarehouseRepository(db)
	warehouseService := NewWarehouseService(db, yearRepo, warehouseRepo)
	warehouseHandler := NewWarehouseHandler(warehouseService)

	WarehouseRoute(app, warehouseHandler)
}
