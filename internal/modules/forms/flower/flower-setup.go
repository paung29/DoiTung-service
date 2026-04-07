package flower

import (
	"github.com/doitung/DoiTung-service/internal/modules/cluster"
	"github.com/doitung/DoiTung-service/internal/modules/year"
	"github.com/doitung/DoiTung-service/internal/modules/zone"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Setup(app *fiber.App, db *gorm.DB) {
	yearRepo := year.NewYearRepository(db)
	zoneRepo := zone.NewZoneRepository(db)
	clusterRepo := cluster.NewClusterRepository(db)
	flowerRepo := NewFlowerRepository(db)
	flowerService := NewFlowerService(db, yearRepo, zoneRepo, clusterRepo, flowerRepo)

	FlowerHandler := NewFlowerHandler(flowerService)
	FlowerRoutes(app, FlowerHandler)
}
