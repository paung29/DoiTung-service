package dashboard

import (
	"strconv"

	"github.com/doitung/DoiTung-service/internal/utils"
	"github.com/gofiber/fiber/v2"
)

type DashboardHandler struct {
	service DashboardService
}

func NewDashboardHandler(service DashboardService) *DashboardHandler {
	return &DashboardHandler{
		service: service,
	}
}

func (h *DashboardHandler) GetPerformanceOverview(context *fiber.Ctx) error {
	yearStr := context.Query("year")
	if yearStr == "" {
		return utils.HandleError(
			context,
			utils.BadRequestError("year is required"),
		)
	}

	year, err := strconv.Atoi(yearStr)
	if err != nil || year <= 0 {
		return utils.HandleError(
			context,
			utils.BadRequestError("year must be a positive integer"),
		)
	}

	response, err := h.service.GetPerformanceOverview(year)
	if err != nil {
		return utils.HandleError(context, err)
	}

	return context.Status(fiber.StatusOK).JSON(response)
}

func (h *DashboardHandler) GetConditionByStage(context *fiber.Ctx) error {
	yearStr := context.Query("year")
	if yearStr == "" {
		return utils.HandleError(
			context,
			utils.BadRequestError("year is required"),
		)
	}

	year, err := strconv.Atoi(yearStr)
	if err != nil || year <= 0 {
		return utils.HandleError(
			context,
			utils.BadRequestError("year must be a positive integer"),
		)
	}

	response, err := h.service.GetConditionByStage(year)
	if err != nil {
		return utils.HandleError(context, err)
	}

	return context.Status(fiber.StatusOK).JSON(response)
}

func (h *DashboardHandler) GetFlowerProductionTrend(context *fiber.Ctx) error {
	response, err := h.service.GetFlowerProductionTrend()
	if err != nil {
		return utils.HandleError(context, err)
	}

	return context.Status(fiber.StatusOK).JSON(response)
}

func (h *DashboardHandler) GetPodProductionTrend(context *fiber.Ctx) error {
	response, err := h.service.GetPodProductionTrend()
	if err != nil {
		return utils.HandleError(context, err)
	}

	return context.Status(fiber.StatusOK).JSON(response)
}

func (h *DashboardHandler) GetPodSetRateTrend(context *fiber.Ctx) error {
	response, err := h.service.GetPodSetRateTrend()
	if err != nil {
		return utils.HandleError(context, err)
	}

	return context.Status(fiber.StatusOK).JSON(response)
}

func (h *DashboardHandler) GetHarvestablePodsTrend(context *fiber.Ctx) error {
	response, err := h.service.GetHarvestablePodsTrend()
	if err != nil {
		return utils.HandleError(context, err)
	}

	return context.Status(fiber.StatusOK).JSON(response)
}

func (h *DashboardHandler) GetFreshPodGradeTrend(context *fiber.Ctx) error {
	response, err := h.service.GetFreshPodGradeTrend()
	if err != nil {
		return utils.HandleError(context, err)
	}

	return context.Status(fiber.StatusOK).JSON(response)
}

func (h *DashboardHandler) GetProductivePolesTrend(context *fiber.Ctx) error {
	response, err := h.service.GetProductivePolesTrend()
	if err != nil {
		return utils.HandleError(context, err)
	}

	return context.Status(fiber.StatusOK).JSON(response)
}
