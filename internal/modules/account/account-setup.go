package account

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)



func Setup(app *fiber.App, db *gorm.DB) {
	accountRepo := NewAccountrepository(db)
	accountService := NewAuthService(accountRepo)
	accountHandler := NewAccountHandler(accountService)

	RegisterRoutes(app, accountHandler)
}