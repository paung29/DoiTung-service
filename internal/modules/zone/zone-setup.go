package zone

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"github.com/doitung/DoiTung-service/internal/modules/year"
)

func Setup(app *fiber.App, db *gorm.DB) {
	zoneRepo := NewZoneRepository(db)
	yearRepo := year.NewYearRepository(db)
	zoneService := NewZoneService(db, zoneRepo, yearRepo)
	zoneHandler := NewZoneHandler(zoneService)

	RegisterRoutes(app, zoneHandler)
}