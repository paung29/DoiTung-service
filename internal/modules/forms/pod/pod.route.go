package pod

import (
	"github.com/doitung/DoiTung-service/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func PodRoutes(app *fiber.App, handler *PodHandler) {
	podGroup := app.Group("pods")

	podGroup.Post("/create", middleware.RequiredAuth, handler.CreateOrUpdatePodForm)
	podGroup.Get("/get-pod-form", middleware.RequiredAuth, handler.GetPodFormDetails)
	podGroup.Get("/get-pod-form-histories", middleware.RequiredAuth, handler.GetPodFormHistories)
	podGroup.Put("/update-pod-form", middleware.RequiredAuth, handler.CreateOrUpdatePodForm)
	podGroup.Get("/get-pod-forms-by-zone", middleware.RequiredAuth, handler.GetPodFormsByZoneId)
}
