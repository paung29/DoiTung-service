package flower

import (
	"github.com/doitung/DoiTung-service/internal/utils"
	"github.com/gofiber/fiber/v2"
)

type FlowerHandler struct {
	service FlowerService
}

func NewFlowerHandler(service FlowerService) *FlowerHandler {
	return &FlowerHandler{service: service}
}

func (h *FlowerHandler) CreateCluster(context *fiber.Ctx) error {

	var userId uint = context.Locals("account_id").(uint)
	var form FlowerFormRequest

	if err := utils.ParseAndValidate(context, &form); err != nil {
		return utils.HandleError(context, err)
	}

	response, err := h.service.CreateFlowerForm(form, userId)

	if err != nil {
		return utils.HandleError(context, err)
	}
	return context.Status(fiber.StatusCreated).JSON(response)
}
