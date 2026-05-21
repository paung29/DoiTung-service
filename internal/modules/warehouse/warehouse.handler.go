package warehouse

import (
	"strconv"

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

func (h *WarehouseHandler) GetWarehouses(context *fiber.Ctx) error {
	response, err := h.service.GetAllWarehouses()

	if err != nil {
		return utils.HandleError(context, err)
	}
	return context.Status(fiber.StatusOK).JSON(response)
}

func (h *WarehouseHandler) GetWarehouseById(context *fiber.Ctx) error {

	warehouseIdStr := context.Query("warehouseId")
	if warehouseIdStr == "" {
		return utils.HandleError(context, utils.ValidationError("warehouseId query parameter is required", nil))
	}

	warehouseId, err := strconv.Atoi(warehouseIdStr)
	if err != nil {
		return utils.HandleError(context, utils.ValidationError("Invalid warehouse ID", nil))
	}

	response, err := h.service.GetWarehouseById(uint(warehouseId))
	if err != nil {
		return utils.HandleError(context, err)
	}

	return context.Status(fiber.StatusOK).JSON(response)
}

func (h *WarehouseHandler) UpdateWarehouse(context *fiber.Ctx) error {
	var form UpdateWarehouseRequest

	if err := utils.ParseAndValidate(context, &form); err != nil {
		return utils.HandleError(context, err)
	}

	response, err := h.service.UpdateWarehouse(form)

	if err != nil {
		return utils.HandleError(context, err)
	}

	return context.Status(fiber.StatusOK).JSON(response)
}
