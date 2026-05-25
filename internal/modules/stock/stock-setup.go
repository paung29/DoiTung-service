package stock

import (
	"github.com/doitung/DoiTung-service/internal/modules/warehouse"
	"github.com/doitung/DoiTung-service/internal/modules/year"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Setup(app *fiber.App, db *gorm.DB) {
	stockRepo := NewStockRepository(db)
	yearRepo := year.NewYearRepository(db)
	warehouseRepo := warehouse.NewWarehouseRepository(db)
	stockService := NewStockService(stockRepo, yearRepo, warehouseRepo)
	stockHandler := NewStockHandler(stockService)

	RegisterRoutes(app, stockHandler)

}
