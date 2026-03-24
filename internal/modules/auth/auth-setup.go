package auth

import (
	"github.com/doitung/DoiTung-service/internal/modules/account"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Setup(app *fiber.App, db *gorm.DB) {
	accountRepo := account.NewAccountrepository(db)
	authService := NewAuthService(accountRepo)
	authHandler := NewAuthHandler(authService)

	RegisterRoutes(app, authHandler)
}