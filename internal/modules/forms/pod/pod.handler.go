package pod

import (
	"strconv"

	"github.com/doitung/DoiTung-service/internal/utils"
	"github.com/gofiber/fiber/v2"
)

type PodHandler struct {
	service PodService
}

func NewPodHandler(service PodService) *PodHandler {
	return &PodHandler{service: service}
}

func (h *PodHandler) CreateOrUpdatePodForm(context *fiber.Ctx) error {
	var userId uint = context.Locals("account_id").(uint)
	var form PodFormRequest

	if err := utils.ParseAndValidate(context, &form); err != nil {
		return utils.HandleError(context, err)
	}

	response, err := h.service.CreateOrUpdatePodForm(form, userId)
	if err != nil {
		return utils.HandleError(context, err)
	}

	return context.Status(fiber.StatusCreated).JSON(response)
}

func (h *PodHandler) GetPodFormDetails(context *fiber.Ctx) error {

	clusterId, err := utils.GetClusterIDFromQuery(context)
	if err != nil {
		return utils.HandleError(context, err)
	}

	response, err := h.service.GetPodFormDetails(uint(clusterId))
	if err != nil {
		return utils.HandleError(context, err)
	}
	return context.JSON(response)
}

func (h *PodHandler) GetPodFormHistories(context *fiber.Ctx) error {
	userId := context.Locals("account_id").(uint)

	yearStr := context.Query("year")
	if yearStr == "" {
		return utils.HandleError(context, utils.BadRequestError("year is required"))
	}

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		return utils.HandleError(context, utils.BadRequestError("invalid year"))
	}

	response, err := h.service.GetPodFormHistories(userId, uint(year))
	if err != nil {
		return utils.HandleError(context, err)
	}
	return context.JSON(response)
}

func (h *PodHandler) GetPodFormsByZoneId(context *fiber.Ctx) error {
	zoneIdStr := context.Query("zoneId")
	if zoneIdStr == "" {
		return utils.HandleError(context, utils.BadRequestError("zoneId is required"))
	}

	zoneId, err := strconv.Atoi(zoneIdStr)
	if err != nil {
		return utils.HandleError(context, utils.BadRequestError("invalid zoneId"))
	}

	if zoneId <= 0 {
		return utils.HandleError(context, utils.BadRequestError("zoneId must be a positive integer"))
	}

	response, err := h.service.GetPodFormsByZoneId(uint(zoneId))
	if err != nil {
		return utils.HandleError(context, err)
	}
	return context.JSON(response)
}
