package stock

import (
	"strconv"

	"github.com/doitung/DoiTung-service/internal/types/enums"
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

func (h *handler) GetCustomerStockTableByYear(context *fiber.Ctx) error {

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
	response, err := h.service.GetCustomerStockTableByYear(uint(year))
	if err != nil {
		return utils.HandleError(context, err)
	}
	return context.Status(fiber.StatusOK).JSON(response)
}

func (h *handler) GetStockOverviewBalanceByYear(context *fiber.Ctx) error {

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
	response, err := h.service.GetStockOverviewBalanceByYear(uint(year))
	if err != nil {
		return utils.HandleError(context, err)
	}
	return context.Status(fiber.StatusOK).JSON(response)
}

func (h *handler) GetStockMovementHistoryFilter(context *fiber.Ctx) error {
	yearStr := context.Query("year")
	categoryStr := context.Query("category")
	gradeStr := context.Query("grade")
	productionYearStr := context.Query("productionYear")
	warehouseIDStr := context.Query("warehouseId")

	if yearStr == "" {
		return utils.HandleError(context, utils.BadRequestError("year is required"))
	}

	year, err := strconv.Atoi(yearStr)
	if err != nil || year <= 0 {
		return utils.HandleError(context, utils.BadRequestError("year must be a positive integer"))
	}

	var category *enums.MovementType
	if categoryStr != "" {
		parsedCategory := enums.MovementType(categoryStr)

		if parsedCategory != enums.MovementCarryOver &&
			parsedCategory != enums.MovementIncoming &&
			parsedCategory != enums.MovementIssued {
			return utils.HandleError(context, utils.BadRequestError("invalid category"))
		}

		category = &parsedCategory
	}

	var grade *enums.Grade
	if gradeStr != "" {
		parsedGrade := enums.Grade(gradeStr)

		if parsedGrade != enums.GradeAPlus &&
			parsedGrade != enums.GradeA &&
			parsedGrade != enums.GradeB &&
			parsedGrade != enums.GradeC &&
			parsedGrade != enums.GradeD &&
			parsedGrade != enums.GradeDPlus {
			return utils.HandleError(context, utils.BadRequestError("invalid grade"))
		}

		grade = &parsedGrade
	}

	var productionYear *uint
	if productionYearStr != "" {
		parsedProductionYear, err := strconv.Atoi(productionYearStr)
		if err != nil || parsedProductionYear <= 0 {
			return utils.HandleError(context, utils.BadRequestError("productionYear must be a positive integer"))
		}

		productionYearUint := uint(parsedProductionYear)
		productionYear = &productionYearUint
	}

	var warehouseID *uint
	if warehouseIDStr != "" {
		parsedWarehouseID, err := strconv.Atoi(warehouseIDStr)
		if err != nil || parsedWarehouseID <= 0 {
			return utils.HandleError(context, utils.BadRequestError("warehouseId must be a positive integer"))
		}

		warehouseIDUint := uint(parsedWarehouseID)
		warehouseID = &warehouseIDUint
	}

	response, err := h.service.GetStockMovementHistoryFilter(
		uint(year),
		category,
		grade,
		productionYear,
		warehouseID,
	)
	if err != nil {
		return utils.HandleError(context, err)
	}

	return context.Status(fiber.StatusOK).JSON(response)
}
