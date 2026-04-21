package pollination

import (
	"github.com/doitung/DoiTung-service/internal/utils"
	"github.com/gofiber/fiber/v2"
)

type PollinationHandler struct {
	service PollinationService
}

func NewPollinationHandler(service PollinationService) *PollinationHandler {
	return &PollinationHandler{service: service}
}

func (h *PollinationHandler) CreateOrUpdatePollinationForm(context *fiber.Ctx) error {

	var userId uint = context.Locals("account_id").(uint)
	var form PollinationFormRequest
	if err := utils.ParseAndValidate(context, &form); err != nil {
		return utils.HandleError(context, err)
	}

	response, err := h.service.CreateOrUpdatePollinationForm(form, userId)
	if err != nil {
		return utils.HandleError(context, err)
	}
	return context.Status(fiber.StatusCreated).JSON(response)
}

func (h *PollinationHandler) GetPollinationFormDetails(context *fiber.Ctx) error {

	clusterId, err := utils.GetClusterIDFromQuery(context)
	if err != nil {
		return utils.HandleError(context, err)
	}

	response, err := h.service.GetPollinationFormDetails(uint(clusterId))
	if err != nil {
		return utils.HandleError(context, err)
	}
	return context.JSON(response)
}
