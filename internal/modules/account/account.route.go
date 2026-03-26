package account

import (
	"github.com/doitung/DoiTung-service/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, handler *AccountHandler) {

	account := app.Group("/accounts")

	account.Post("/create",middleware.RequiredAuth, middleware.RequireRoles("ADMIN"), handler.CreateAccount)

}