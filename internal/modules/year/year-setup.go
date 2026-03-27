package year

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Setup(app *fiber.App, db *gorm.DB) {
	yearRepo := NewYearRepository(db)
	yearService := NewYearService(db, yearRepo)
	yearHandler := NewYearHandler(yearService)

	RegisterRoutes(app, yearHandler)
}