package auth

import (
	"github.com/doitung/DoiTung-service/internal/utils"
	"github.com/gofiber/fiber/v2"
)

func Login(context *fiber.Ctx) error {

	var form LoginRequest

	if err := context.BodyParser(&form);err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success" : false,
			"message" : "invalid json format",
		})
	}

	if err := utils.Validate.Struct(form);err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success" : false,
			"message" : "validation error",
			"errors" : utils.FormatValidationErrors(err),
		})
	}

	return context.JSON(fiber.Map{
		"success" : true,
	})

}