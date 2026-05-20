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
