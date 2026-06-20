package exportdata

import (
	"fmt"
	"strconv"

	"github.com/doitung/DoiTung-service/internal/utils"
	"github.com/gofiber/fiber/v2"
)

type ExportDataHandler struct {
	service ExportDataService
}

func NewExportDataHandler(service ExportDataService) *ExportDataHandler {
	return &ExportDataHandler{
		service: service,
	}
}

func sendExcel(
	context *fiber.Ctx,
	response ExportXLSXResponse,
) error {
	context.Set(
		fiber.HeaderContentType,
		"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
	)
	context.Set(
		fiber.HeaderContentDisposition,
		fmt.Sprintf(
			`attachment; filename="%s"`,
			response.FileName,
		),
	)
	context.Set(
		fiber.HeaderContentLength,
		strconv.Itoa(len(response.FileBytes)),
	)
	context.Set("X-Content-Type-Options", "nosniff")

	return context.Status(fiber.StatusOK).
		Send(response.FileBytes)
}

func (h *ExportDataHandler) ExportClusterFormsXLSX(context *fiber.Ctx) error {
	yearStr := context.Query("year")
	if yearStr == "" {
		return utils.HandleError(context, utils.BadRequestError("year is required"))
	}

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		return utils.HandleError(context, utils.BadRequestError("invalid year"))
	}
	if year <= 0 {
		return utils.HandleError(context, utils.BadRequestError("year must be a positive integer"))
	}

	response, err := h.service.ExportClusterFormsXLSX(uint(year))
	if err != nil {
		return utils.HandleError(context, err)
	}

	return sendExcel(context, response)
}

func (h *ExportDataHandler) ExportHarvestGrading(context *fiber.Ctx) error {
	yearStr := context.Query("year")
	if yearStr == "" {
		return utils.HandleError(context, utils.BadRequestError("year is required"))
	}

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		return utils.HandleError(context, utils.BadRequestError("invalid year"))
	}
	if year <= 0 {
		return utils.HandleError(context, utils.BadRequestError("year must be a positive integer"))
	}

	response, err := h.service.ExportHarvestGrading(year)
	if err != nil {
		return utils.HandleError(context, err)
	}

	return sendExcel(context, response)
}

func (h *ExportDataHandler) ExportHarvestGradingSummary(context *fiber.Ctx) error {
	yearStr := context.Query("year")
	if yearStr == "" {
		return utils.HandleError(context, utils.BadRequestError("year is required"))
	}

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		return utils.HandleError(context, utils.BadRequestError("invalid year"))
	}
	if year <= 0 {
		return utils.HandleError(context, utils.BadRequestError("year must be a positive integer"))
	}

	response, err := h.service.ExportHarvestGradingSummary(year)
	if err != nil {
		return utils.HandleError(context, err)
	}

	return sendExcel(context, response)
}

func (h *ExportDataHandler) ExportStockMovements(context *fiber.Ctx) error {
	yearStr := context.Query("year")
	if yearStr == "" {
		return utils.HandleError(context, utils.BadRequestError("year is required"))
	}

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		return utils.HandleError(context, utils.BadRequestError("invalid year"))
	}
	if year <= 0 {
		return utils.HandleError(context, utils.BadRequestError("year must be a positive integer"))
	}
	response, err := h.service.ExportStockMovements(year)
	if err != nil {
		return utils.HandleError(context, err)
	}

	return sendExcel(context, response)

}
