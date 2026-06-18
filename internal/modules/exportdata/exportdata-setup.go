package exportdata

import (
	"github.com/doitung/DoiTung-service/internal/modules/year"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Setup(app *fiber.App, db *gorm.DB) {
	exportDataRepo := NewExportDataRepository(db)
	yearRepo := year.NewYearRepository(db)

	exportDataService := NewExportDataService(yearRepo, exportDataRepo)
	exportDataHandler := NewExportDataHandler(exportDataService)

	RegisterRoutes(app, exportDataHandler)
}
