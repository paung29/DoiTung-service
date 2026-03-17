package auth

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(app *fiber.App) {

	auth := app.Group("/auth")

	auth.Post("/login", Login)

}