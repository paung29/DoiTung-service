package middleware

import (
	"github.com/doitung/DoiTung-service/internal/utils"
	"github.com/gofiber/fiber/v2"
)

func RequireRoles(allowedRoles ...string) fiber.Handler {

	return func(context *fiber.Ctx) error {
		
		role, ok := context.Locals("role").(string)
		if !ok || role == "" {
			return utils.UnauthorizedError("role not found")
		}

		for _, allowed := range allowedRoles {
			if role == allowed {
				return context.Next()
			}
		}

		return utils.ForbiddenError("Not Allowed Access")
	}
}