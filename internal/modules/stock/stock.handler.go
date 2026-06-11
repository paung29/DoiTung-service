package stock

import (
	"strconv"

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
	var form CreateCarryOverStockRequest

	if err := utils.ParseAndValidate(context, &form); err != nil {
		return utils.HandleError(context, err)
	}
	response, err := h.service.CreateCarryOver(accountID, form)
	if err != nil {
		return utils.HandleError(context, err)
	}

	return context.Status(fiber.StatusOK).JSON(response)
}

func (h *handler) CreateIncomingStock(context *fiber.Ctx) error {
	var accountID uint = context.Locals("account_id").(uint)
	var form CreateIncomingStockRequest

	if err := utils.ParseAndValidate(context, &form); err != nil {
		return utils.HandleError(context, err)
	}
	response, err := h.service.CreateIncomingStock(accountID, form)
	if err != nil {
		return utils.HandleError(context, err)
	}

	return context.Status(fiber.StatusOK).JSON(response)
}

func (h *handler) CreateIssuedStock(context *fiber.Ctx) error {
	var accountID uint = context.Locals("account_id").(uint)
	var form CreateIssuedStockRequest

	if err := utils.ParseAndValidate(context, &form); err != nil {
		return utils.HandleError(context, err)
	}
	response, err := h.service.CreateIssuedStock(accountID, form)
	if err != nil {
		return utils.HandleError(context, err)
	}

	return context.Status(fiber.StatusOK).JSON(response)
}

// func (h *handler) UpdateStockMovement(context *fiber.Ctx) error {
// 	var accountID uint = context.Locals("account_id").(uint)
// 	var form UpdateStockMovementRequest

// 	if err := utils.ParseAndValidate(context, &form); err != nil {
// 		return utils.HandleError(context, err)
// 	}
// 	response, err := h.service.UpdateStockMovement(accountID, form)
// 	if err != nil {
// 		return utils.HandleError(context, err)
// 	}

// 	return context.Status(fiber.StatusOK).JSON(response)
// }

func (h *handler) DeleteStockMovement(context *fiber.Ctx) error {

	stockMovementIdStr := context.Query("stock_movement_id")
	if stockMovementIdStr == "" {
		return utils.HandleError(context, utils.BadRequestError("Missing stock_movement_id query parameter"))
	}

	stockMovementId, err := strconv.Atoi(stockMovementIdStr)
	if err != nil {
		return utils.HandleError(context, utils.BadRequestError("Invalid stock_movement_id query parameter"))
	}

	if stockMovementId <= 0 {
		return utils.HandleError(context, utils.BadRequestError("stock_movement_id must be a positive integer"))
	}

	response, err := h.service.DeleteStockMovement(uint(stockMovementId))
	if err != nil {
		return utils.HandleError(context, err)
	}

	return context.Status(fiber.StatusOK).JSON(response)
}

func (h *handler) GetStockMovementListsByYear(context *fiber.Ctx) error {

	yearStr := context.Query("year")
	if yearStr == "" {
		return utils.HandleError(context, utils.BadRequestError("Missing year query parameter"))
	}

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		return utils.HandleError(context, utils.BadRequestError("Invalid year query parameter"))
	}

	if year <= 0 {
		return utils.HandleError(context, utils.BadRequestError("year must be a positive integer"))
	}

	response, err := h.service.GetStockMovementListsByYear(uint(year))
	if err != nil {
		return utils.HandleError(context, err)
	}

	return context.Status(fiber.StatusOK).JSON(response)
}
