package warehouse

import (
	"github.com/doitung/DoiTung-service/internal/utils"
	"github.com/gofiber/fiber/v2"
)

type WarehouseHandler struct {
	service WarehouseService
}

func NewWarehouseHandler(service WarehouseService) *WarehouseHandler {
	return &WarehouseHandler{
		service: service,
	}
}

func (h *WarehouseHandler) CreateWarehouse(context *fiber.Ctx) error {
	var form CreateWarehouseRequest

	if err := utils.ParseAndValidate(context, &form); err != nil {
		return utils.HandleError(context, err)
	}

	response, err := h.service.CreateWarehouse(form)

	if err != nil {
		return utils.HandleError(context, err)
	}

	return context.Status(fiber.StatusCreated).JSON(response)
}
