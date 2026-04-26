package pod

import (
	"github.com/doitung/DoiTung-service/internal/modules/cluster"
	"github.com/doitung/DoiTung-service/internal/modules/forms/pollination"
	"github.com/doitung/DoiTung-service/internal/modules/year"
	"github.com/doitung/DoiTung-service/internal/modules/zone"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Setup(app *fiber.App, db *gorm.DB) {
	yearRepo := year.NewYearRepository(db)
	zoneRepo := zone.NewZoneRepository(db)
	clusterRepo := cluster.NewClusterRepository(db)
	pollinationRepo := pollination.NewPollinationRepository(db)
	podRepo := NewPodRepository(db)
	podService := NewPodService(db, yearRepo, zoneRepo, clusterRepo, pollinationRepo, podRepo)

	podHandler := NewPodHandler(podService)
	PodRoutes(app, podHandler)
}
