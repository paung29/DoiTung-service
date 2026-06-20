package exportdata

import (
	"github.com/doitung/DoiTung-service/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, handler *ExportDataHandler) {
	exportData := app.Group("/export-data")

	exportData.Get(
		"/cluster-forms",
		middleware.RequiredAuth,
		middleware.RequireRoles("ADMIN"),
		handler.ExportClusterFormsXLSX,
	)
	exportData.Get(
		"/harvest-grading",
		middleware.RequiredAuth,
		middleware.RequireRoles("ADMIN"),
		handler.ExportHarvestGrading,
	)
}
