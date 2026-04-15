package harvestgrading

import (
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
