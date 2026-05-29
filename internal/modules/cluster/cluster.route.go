package cluster

import (
	"github.com/doitung/DoiTung-service/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func ClusterRoutes(app *fiber.App, handler *ClusterHandler) {
	cluster := app.Group("/clusters")

	cluster.Post("/create", middleware.RequiredAuth, handler.CreateCluster)
	cluster.Get("/get-by-zone", middleware.RequiredAuth, handler.GetClustersByZone)

	cluster.Get("/get-cluster-form", middleware.RequiredAuth, handler.GetClusterForm)
	cluster.Put("/update-cluster-form", middleware.RequiredAuth, handler.UpdateClusterForm)
	cluster.Get("/get-cluster-form-histories", middleware.RequiredAuth, handler.GetClusterFormHistories)
	cluster.Get("/get-all-cluster-form-by-zone", middleware.RequiredAuth, middleware.RequireRoles("ADMIN"), handler.GetAllClustersFormByZone)
}
