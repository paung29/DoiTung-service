package preharvest

import (
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

	response, err := h.service.CreateOreUpdatePreHarvestForm(form, userId)
	if err != nil {
		return utils.HandleError(context, err)
	}

	return context.Status(fiber.StatusCreated).JSON(response)
}
