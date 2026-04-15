package harvestgrading

import (
	"github.com/doitung/DoiTung-service/internal/modules/pole"
	"github.com/doitung/DoiTung-service/internal/modules/year"
	"github.com/doitung/DoiTung-service/internal/modules/zone"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Setup(app *fiber.App, db *gorm.DB) {
	yearRepo := year.NewYearRepository(db)
	zoneRepo := zone.NewZoneRepository(db)
	poleRepo := pole.NewPoleRepository(db)
	harvestGradingRepo := NewHarvestGradingRepository(db)

	harvestGradingService := NewHarvestGradingService(db, yearRepo, zoneRepo, poleRepo, harvestGradingRepo)
	harvestGradingHandler := NewHarvestGradingHandler(harvestGradingService)

	HarvestGradingRoute(app, harvestGradingHandler)

}
