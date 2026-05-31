package pollination

import (
	"strconv"

	"github.com/doitung/DoiTung-service/internal/utils"
	"github.com/gofiber/fiber/v2"
)

type PollinationHandler struct {
	service PollinationService
}

func NewPollinationHandler(service PollinationService) *PollinationHandler {
	return &PollinationHandler{service: service}
}

func (h *PollinationHandler) CreateOrUpdatePollinationForm(context *fiber.Ctx) error {

	var userId uint = context.Locals("account_id").(uint)
	var form PollinationFormRequest
	if err := utils.ParseAndValidate(context, &form); err != nil {
		return utils.HandleError(context, err)
	}

	response, err := h.service.CreateOrUpdatePollinationForm(form, userId)
	if err != nil {
		return utils.HandleError(context, err)
	}
	return context.Status(fiber.StatusCreated).JSON(response)
}

func (h *PollinationHandler) GetPollinationFormDetails(context *fiber.Ctx) error {

	clusterId, err := utils.GetClusterIDFromQuery(context)
	if err != nil {
		return utils.HandleError(context, err)
	}

	response, err := h.service.GetPollinationFormDetails(uint(clusterId))
	if err != nil {
		return utils.HandleError(context, err)
	}
	return context.JSON(response)
}

func (h *PollinationHandler) GetPollinationFormHistories(context *fiber.Ctx) error {

	userId := context.Locals("account_id").(uint)
	yearStr := context.Query("year")
	if yearStr == "" {
		return utils.HandleError(context, utils.BadRequestError("year is required"))
	}

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		return utils.HandleError(context, utils.BadRequestError("invalid year"))
	}

	response, err := h.service.GetPollinationFormHistories(userId, uint(year))
	if err != nil {
		return utils.HandleError(context, err)
	}
	return context.JSON(response)
}

func (h *PollinationHandler) GetPollinationFormsByZoneId(context *fiber.Ctx) error {

	zoneIdStr := context.Query("zoneId")
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

	response, err := h.service.GetPollinationFormsByZoneId(uint(zoneId))
	if err != nil {
		return utils.HandleError(context, err)
	}
	return context.JSON(response)
}
