package auth

import (
	"github.com/doitung/DoiTung-service/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, handler *AuthHandler) {

	auth := app.Group("/auth")

	auth.Post("/login", handler.Login)

	auth.Get("/me", middleware.RequiredAuth, middleware.RequireRoles("ADMIN", "STAFF"),handler.GetUserInfo)

}