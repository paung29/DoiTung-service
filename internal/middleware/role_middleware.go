package middleware

import "github.com/gofiber/fiber/v2"

func RequireRoles(allowedRoles ...string) fiber.Handler {

	return func(context *fiber.Ctx) error {
		
		role, ok := context.Locals("role").(string)
		if !ok || role == "" {
			return context.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"success": false,
				"message": "role not found",
			})
		}

		for _, allowed := range allowedRoles {
			if role == allowed {
				return context.Next()
			}
		}

		return context.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"success": false,
			"message": "forbidden",
		})
	}
}