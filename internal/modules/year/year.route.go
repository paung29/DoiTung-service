package year

import (
	"github.com/doitung/DoiTung-service/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, handler *YearHandler) {

	account := app.Group("/years")

	account.Post("/create",middleware.RequiredAuth, middleware.RequireRoles("ADMIN"), handler.CreateYear)
	account.Put("/form-setting/update",middleware.RequiredAuth, middleware.RequireRoles("ADMIN"), handler.ChangeYearFormSettingStatus)
    account.Get("/get-all-years", middleware.RequiredAuth, middleware.RequireRoles("STAFF","ADMIN"), handler.GetYears)
}