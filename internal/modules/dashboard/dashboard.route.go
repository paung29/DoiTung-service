package dashboard

import (
	"github.com/doitung/DoiTung-service/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, handler *DashboardHandler) {
	dashboard := app.Group("/dashboard")

	dashboard.Get(
		"/performance-overview",
		middleware.RequiredAuth,
		middleware.RequireRoles("ADMIN"),
		handler.GetPerformanceOverview,
	)
}
