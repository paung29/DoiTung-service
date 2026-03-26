package account

import (
	"github.com/doitung/DoiTung-service/internal/utils"
	"github.com/gofiber/fiber/v2"
)

type AccountHandler struct {
	service AccountService
}

func NewAccountHandler(service AccountService) *AccountHandler {
	return &AccountHandler{
		service: service,
	}
}

func (h *AccountHandler) CreateAccount(context *fiber.Ctx) error {
	var form AccountCreateForm

	if err := context.BodyParser(&form); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "invalid json format",
		})
	}

	if err := utils.Validate.Struct(form); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "validation error",
			"errors":  utils.FormatValidationErrors(err),
		})
	}

	response, err := h.service.CreateAccount(form)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "internal server error",
		})
	}
	
	if !response.Success {
		return context.Status(fiber.StatusBadRequest).JSON(response)
	}
	
	return context.Status(fiber.StatusCreated).JSON(response)
}