package pole

import (
	"github.com/doitung/DoiTung-service/internal/modules/year"
	"github.com/doitung/DoiTung-service/internal/modules/zone"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Setup(app *fiber.App, db *gorm.DB) {
	yearRepo := year.NewYearRepository(db)
	zoneRepo := zone.NewZoneRepository(db)
	poleRepo := NewPoleRepository(db)

	poleService := NewPoleService(db, yearRepo, zoneRepo, poleRepo)
	poleHandler := NewPoleHandler(poleService)

	PoleRoute(app, poleHandler)
}
