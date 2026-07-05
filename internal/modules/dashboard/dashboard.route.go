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
	dashboard.Get(
		"/condition-by-stage",
		middleware.RequiredAuth,
		middleware.RequireRoles("ADMIN"),
		handler.GetConditionByStage,
	)

	dashboard.Get(
		"/flower-production-trend",
		middleware.RequiredAuth,
		middleware.RequireRoles("ADMIN"),
		handler.GetFlowerProductionTrend,
	)

	dashboard.Get(
		"/pod-production-trend",
		middleware.RequiredAuth,
		middleware.RequireRoles("ADMIN"),
		handler.GetPodProductionTrend,
	)

	dashboard.Get(
		"/pod-set-rate-trend",
		middleware.RequiredAuth,
		middleware.RequireRoles("ADMIN"),
		handler.GetPodSetRateTrend,
	)

	dashboard.Get(
		"/harvestable-pods-trend",
		middleware.RequiredAuth,
		middleware.RequireRoles("ADMIN"),
		handler.GetHarvestablePodsTrend,
	)

	dashboard.Get(
		"/fresh-pod-grade-trend",
		middleware.RequiredAuth,
		middleware.RequireRoles("ADMIN"),
		handler.GetFreshPodGradeTrend,
	)
}
