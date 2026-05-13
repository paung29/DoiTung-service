package preharvest

import (
	"strconv"

	"github.com/doitung/DoiTung-service/internal/utils"
	"github.com/gofiber/fiber/v2"
)

type preHarvestHandler struct {
	service PreHarvestService
}

func NewPreHarvestHandler(service PreHarvestService) *preHarvestHandler {
	return &preHarvestHandler{service: service}
}

func (h *preHarvestHandler) CreateOrUpdatePreHarvestForm(context *fiber.Ctx) error {
	var userId uint = context.Locals("account_id").(uint)
	var form PreHarvestFormRequest

	if err := utils.ParseAndValidate(context, &form); err != nil {
		return utils.HandleError(context, err)
	}

	response, err := h.service.CreateOrUpdatePreHarvestForm(form, userId)
	if err != nil {
		return utils.HandleError(context, err)
	}

	return context.Status(fiber.StatusCreated).JSON(response)
}

func (h *preHarvestHandler) GetPreHarvestFormDetails(context *fiber.Ctx) error {

	clusterId, err := utils.GetClusterIDFromQuery(context)
	if err != nil {
		return utils.HandleError(context, err)
	}

	response, err := h.service.GetPreHarvestFormDetails(uint(clusterId))
	if err != nil {
		return utils.HandleError(context, err)
	}
	return context.JSON(response)
}

func (h *preHarvestHandler) GetPreHarvestFormHistories(context *fiber.Ctx) error {
	userId := context.Locals("account_id").(uint)

	yearStr := context.Query("year")
	if yearStr == "" {
		return utils.HandleError(context, utils.BadRequestError("year query parameter is required"))
	}

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		return utils.HandleError(context, utils.BadRequestError("invalid year format"))
	}

	response, err := h.service.GetPreHarvestFormHistories(userId, uint(year))
	if err != nil {
		return utils.HandleError(context, err)
	}
	return context.JSON(response)
}
