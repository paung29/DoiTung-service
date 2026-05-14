package harvestgrading

import (
	"strconv"

	"github.com/doitung/DoiTung-service/internal/utils"
	"github.com/gofiber/fiber/v2"
)

type HarvestGradingHandler struct {
	service HarvestGradingService
}

func NewHarvestGradingHandler(service HarvestGradingService) *HarvestGradingHandler {
	return &HarvestGradingHandler{
		service: service,
	}
}

func (h *HarvestGradingHandler) CreateOrUpdateHarvestGradingForm(context *fiber.Ctx) error {

	var userId uint = context.Locals("account_id").(uint)
	var form HarvestGradingFormRequest

	if err := utils.ParseAndValidate(context, &form); err != nil {
		return utils.HandleError(context, err)
	}

	response, err := h.service.CreateOrUpdateHarvestGradingForm(form, userId)
	if err != nil {
		return utils.HandleError(context, err)
	}

	return context.Status(fiber.StatusCreated).JSON(response)
}

func (h *HarvestGradingHandler) GetHarvestGradingFormDetails(context *fiber.Ctx) error {

	poleIdStr := context.Query("poleId")
	if poleIdStr == "" {
		return utils.HandleError(context, utils.BadRequestError("poleId is required"))
	}

	poleId, err := strconv.Atoi(poleIdStr)
	if err != nil {
		return utils.HandleError(context, utils.BadRequestError("invalid poleId"))
	}

	if poleId <= 0 {
		return utils.HandleError(context, utils.BadRequestError("poleId must be a positive integer"))
	}

	response, err := h.service.GetHarvestGradingFormDetailsByPoleID(uint(poleId))

	if err != nil {
		return utils.HandleError(context, err)
	}
	return context.JSON(response)
}

func (h *HarvestGradingHandler) GetHarvestGradingFormHistories(context *fiber.Ctx) error {

	userId := context.Locals("account_id").(uint)

	yearStr := context.Query("year")
	if yearStr == "" {
		return utils.HandleError(context, utils.BadRequestError("year is required"))
	}

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		return utils.HandleError(context, utils.BadRequestError("invalid year"))
	}

	if year <= 0 {
		return utils.HandleError(context, utils.BadRequestError("year must be a positive integer"))
	}

	response, err := h.service.GetHarvestGradingFormHistories(userId, uint(year))
	if err != nil {
		return utils.HandleError(context, err)
	}
	return context.JSON(response)
}
