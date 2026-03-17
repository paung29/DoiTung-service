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

	token, response, err := LoginService(form)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "internal server error",
		})
	}

	if !response.Success {
		return context.Status(fiber.StatusUnauthorized).JSON(response)
	}

	context.Cookie(&fiber.Cookie{
		Name : "access_token",
		Value: token,
		HTTPOnly: true,
		Secure: false,
		SameSite: fiber.CookieSameSiteLaxMode,
		Path:     "/",
		MaxAge:   60 * 60 * 24,
	})

	return context.JSON(fiber.Map{
		"success" : true,
	})

}