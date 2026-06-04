package year

import (
	"github.com/doitung/DoiTung-service/internal/utils"
	"github.com/gofiber/fiber/v2"
)

type YearHandler struct {
	service YearService
}

func NewYearHandler(service YearService) *YearHandler {
	return &YearHandler{
		service: service,
	}
}

func (h YearHandler) CreateYear(context *fiber.Ctx) error {
	var form YearCreateForm

	if err := utils.ParseAndValidate(context, &form); err != nil {
		return utils.HandleError(context, err)
	}

	response, err := h.service.CreateYear(form)

	if err != nil {
		return utils.HandleError(context, err)
	}
	return context.Status(fiber.StatusCreated).JSON(response)
}

func (h YearHandler) ChangeYearFormSettingStatus(context *fiber.Ctx) error {
	var form YearFormSettingStatusChange

	if err := utils.ParseAndValidate(context, &form); err != nil {
		return utils.HandleError(context, err)
	}

	response, err := h.service.ChangeYearFormSettingStatus(form)
	if err != nil {
		return utils.HandleError(context, err)
	}
	return context.Status(fiber.StatusOK).JSON(response)
}

func (h YearHandler) GetYears(context *fiber.Ctx) error {

	response, err := h.service.GetYear()
	if err != nil {
		return utils.HandleError(context, err)
	}

	return context.Status(fiber.StatusOK).JSON(response)
}

func (h YearHandler) GetYearDetails(context *fiber.Ctx) error {

	response, err := h.service.GetYearDetails()
	if err != nil {
		return utils.HandleError(context, err)
	}

	return context.Status(fiber.StatusOK).JSON(response)
}
