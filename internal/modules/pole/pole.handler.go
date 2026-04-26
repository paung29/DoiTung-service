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

	zoneNoStr := context.Query("zoneNo")
	if zoneNoStr == "" {
		return utils.HandleError(context, utils.BadRequestError("zoneNo is required"))
	}

	zoneNo, err := strconv.Atoi(zoneNoStr)
	if err != nil {
		return utils.HandleError(context, utils.BadRequestError("invalid zoneNo"))
	}

	response, err := h.service.GetPoleByZone(year, zoneNo)

	if err != nil {
		return utils.HandleError(context, err)
	}

	return context.Status(fiber.StatusOK).JSON(response)
}
