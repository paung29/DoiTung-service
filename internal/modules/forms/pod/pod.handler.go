package pod

import (
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
