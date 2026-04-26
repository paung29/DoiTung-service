package cluster

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
	clusterRepo := NewClusterRepository(db)
	poleRepo := pole.NewPoleRepository(db)
	clusterService := NewClusterService(db, yearRepo, zoneRepo, poleRepo, clusterRepo)
	clusterHandler := NewClusterHandler(clusterService)

	ClusterRoutes(app, clusterHandler)
}
