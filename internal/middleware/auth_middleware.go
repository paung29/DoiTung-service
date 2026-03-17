package middleware

import (
	"github.com/doitung/DoiTung-service/internal/modules/auth"
	"github.com/gofiber/fiber/v2"
)


func RequiredAuth(context *fiber.Ctx) error {

	token := context.Cookies("access_token")

	if token == "" {
		return context.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "authentication required",
		})
	}

	accountID, err := auth.ParseToken(token)

	if err != nil {
		return context.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "invalid or expired token",
		})
	}

	context.Locals("account_id", accountID)

	return context.Next()
}