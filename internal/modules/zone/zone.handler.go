package zone

import (
	"strconv"

	"github.com/doitung/DoiTung-service/internal/utils"
	"github.com/gofiber/fiber/v2"
)

type ZoneHandler struct {
	service ZoneService
}

func NewZoneHandler(service ZoneService) *ZoneHandler {
	return &ZoneHandler{
		service: service,
	}
}

func (h ZoneHandler) CreateZone(context *fiber.Ctx) error {
	var form CreateZoneRequest

	if err := utils.ParseAndValidate(context, &form); err != nil {
		return utils.HandleError(context, err)
	}

	response, err := h.service.CreateZone(form)
	if err != nil {
		return utils.HandleError(context, err)
	}

	return context.Status(fiber.StatusCreated).JSON(response)
}

func (h ZoneHandler) GetAllZone(context *fiber.Ctx) error {

	yearStr := context.Query("year")
	if yearStr == "" {
		return utils.HandleError(context, utils.BadRequestError("year is required"))
	}

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		return utils.HandleError(context, utils.BadRequestError("invalid year"))
	}

	response, err := h.service.GetAllZone(uint(year))
	if err != nil {
		return utils.HandleError(context, err)
	}

	return context.Status(fiber.StatusCreated).JSON(response)
}

func (h ZoneHandler) GetZoneManagementTable(context *fiber.Ctx) error {

	yearStr := context.Query("year")
	if yearStr == "" {
		return utils.HandleError(context, utils.BadRequestError("year is required"))
	}

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		return utils.HandleError(context, utils.BadRequestError("invalid year"))
	}

	response, err := h.service.GetZoneManagementTable(uint(year))
	if err != nil {
		return utils.HandleError(context, err)
	}

	return context.Status(fiber.StatusOK).JSON(response)
}
