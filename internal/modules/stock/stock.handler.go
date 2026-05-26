package stock

import (
	"github.com/doitung/DoiTung-service/internal/utils"
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	service StockService
}

func NewStockHandler(service StockService) *handler {
	return &handler{service: service}
}

func (h *handler) CreateCarryOver(context *fiber.Ctx) error {
	var accountID uint = context.Locals("account_id").(uint)
	var form CreateCarryOverRequest

	if err := utils.ParseAndValidate(context, &form); err != nil {
		return utils.HandleError(context, err)
	}
	response, err := h.service.CreateCarryOver(accountID, form)
	if err != nil {
		return utils.HandleError(context, err)
	}

	return context.Status(fiber.StatusOK).JSON(response)
}
