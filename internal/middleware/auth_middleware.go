package middleware

import (
	"github.com/gofiber/fiber/v2"
	jwtService "github.com/doitung/DoiTung-service/internal/common/jwt"
)


func RequiredAuth(context *fiber.Ctx) error {

	token := context.Cookies("access_token")

	if token == "" {
		return context.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "authentication required",
		})
	}

	claims, err := jwtService.ParseToken(token)

	if err != nil {
		return context.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "invalid or expired token",
		})
	}

	context.Locals("account_id", claims.AccountID)
	context.Locals("role", claims.Role)

	return context.Next()
}