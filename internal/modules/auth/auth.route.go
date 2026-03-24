package auth

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(app *fiber.App, handler *AuthHandler) {

	auth := app.Group("/auth")

	auth.Post("/login", handler.Login)

}