package preharvest

import (
	"github.com/doitung/DoiTung-service/internal/modules/cluster"
	"github.com/doitung/DoiTung-service/internal/modules/forms/pod"
	"github.com/doitung/DoiTung-service/internal/modules/year"
	"github.com/doitung/DoiTung-service/internal/modules/zone"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Setup(app *fiber.App, db *gorm.DB) {
	yearRepo := year.NewYearRepository(db)
	zoneRepo := zone.NewZoneRepository(db)
	clusterRepo := cluster.NewClusterRepository(db)
	podRepo := pod.NewPodRepository(db)
	preHarvestRepo := NewPreHarvestRepository(db)

	preHarvestService := NewPreHarvestService(db, yearRepo, zoneRepo, clusterRepo, podRepo, preHarvestRepo)
	handler := NewPreHarvestHandler(preHarvestService)

	PreHarvestRoute(app, handler)
}
