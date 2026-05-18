package account

import (
	"github.com/doitung/DoiTung-service/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, handler *AccountHandler) {

	account := app.Group("/accounts")

	account.Post("/create", middleware.RequiredAuth, middleware.RequireRoles("ADMIN"), handler.CreateAccount)
	account.Put("/update-info", middleware.RequiredAuth, middleware.RequireRoles("ADMIN"), handler.UpdateAccountInfo)
	account.Put("/update-password", middleware.RequiredAuth, middleware.RequireRoles("ADMIN"), handler.UpdateAccountPassword)
	account.Get("/get-all", middleware.RequiredAuth, middleware.RequireRoles("ADMIN"), handler.GetAllAccounts)
	account.Get("/get-by-id", middleware.RequiredAuth, middleware.RequireRoles("ADMIN"), handler.GetAccountById)
}
