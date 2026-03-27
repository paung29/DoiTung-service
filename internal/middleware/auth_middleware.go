package middleware

import (
	jwtService "github.com/doitung/DoiTung-service/internal/common/jwt"
	"github.com/doitung/DoiTung-service/internal/utils"
	"github.com/gofiber/fiber/v2"
)


func RequiredAuth(context *fiber.Ctx) error {

	token := context.Cookies("access_token")

	if token == "" {
		return utils.UnauthorizedError("authentication required")
	}

	claims, err := jwtService.ParseToken(token)

	if err != nil {
		return utils.UnauthorizedError("invalid or expired token")
	}
		
	context.Locals("account_id", claims.AccountID)
	context.Locals("role", claims.Role)

	return context.Next()
}