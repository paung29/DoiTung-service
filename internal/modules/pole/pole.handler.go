package pole

import (
	"strconv"

	"github.com/doitung/DoiTung-service/internal/utils"
	"github.com/gofiber/fiber/v2"
)

type PoleHandler struct {
	service PoleService
}

func NewPoleHandler(service PoleService) *PoleHandler {
	return &PoleHandler{service: service}
}

func (h *PoleHandler) GetPoleByZone(context *fiber.Ctx) error {

	yearStr := context.Query("year")
	if yearStr == "" {
		return utils.HandleError(context, utils.BadRequestError("year is required"))
	}

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		return utils.HandleError(context, utils.BadRequestError("invalid year"))
	}

	zoneIdStr := context.Query("zoneId")
	if zoneIdStr == "" {
		return utils.HandleError(context, utils.BadRequestError("zoneId is required"))
	}

	zoneId, err := strconv.Atoi(zoneIdStr)
	if err != nil {
		return utils.HandleError(context, utils.BadRequestError("invalid zoneId"))
	}

	response, err := h.service.GetPoleByZone(year, zoneId)

	if err != nil {
		return utils.HandleError(context, err)
	}

	return context.Status(fiber.StatusOK).JSON(response)
}

func (h *PoleHandler) GetPoleFilter(context *fiber.Ctx) error {
	zoneIdStr := context.Query("zoneId")
	poleNoStr := context.Query("poleNo")
	harvestGradingFormDoneStr := context.Query("harvestGradingFormDone")

	if zoneIdStr == "" {
		return utils.HandleError(context, utils.BadRequestError("zoneId is required"))
	}

	zoneId, err := strconv.Atoi(zoneIdStr)
	if err != nil {
		return utils.HandleError(context, utils.BadRequestError("invalid zoneId"))
	}

	if zoneId <= 0 {
		return utils.HandleError(context, utils.BadRequestError("zoneId must be greater than 0"))
	}

	var poleNo *uint
	if poleNoStr != "" {
		parsedPoleNo, err := strconv.Atoi(poleNoStr)
		if err != nil {
			return utils.HandleError(context, utils.BadRequestError("invalid poleNo"))
		}

		if parsedPoleNo <= 0 {
			return utils.HandleError(context, utils.BadRequestError("poleNo must be greater than 0"))
		}

		poleNoUint := uint(parsedPoleNo)
		poleNo = &poleNoUint
	}

	var harvestGradingFormDone *bool
	if harvestGradingFormDoneStr != "" {
		parsedHarvestGradingFormDone, err := strconv.ParseBool(harvestGradingFormDoneStr)
		if err != nil {
			return utils.HandleError(context, utils.BadRequestError("invalid harvestGradingFormDone"))
		}

		harvestGradingFormDone = &parsedHarvestGradingFormDone
	}

	response, err := h.service.GetPoleFilter(
		uint(zoneId),
		poleNo,
		harvestGradingFormDone,
	)
	if err != nil {
		return utils.HandleError(context, err)
	}

	return context.Status(fiber.StatusOK).JSON(response)
}
