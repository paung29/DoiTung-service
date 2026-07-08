package dashboard

import (
	"github.com/doitung/DoiTung-service/internal/modules/year"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Setup(app *fiber.App, db *gorm.DB) {
	yearRepo := year.NewYearRepository(db)
	dashboardRepo := NewDashboardRepository(db)
	dashboardService := NewDashboardService(yearRepo, dashboardRepo)
	dashboardHandler := NewDashboardHandler(dashboardService)

	RegisterRoutes(app, dashboardHandler)
}
