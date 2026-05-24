package customer

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Setup(app *fiber.App, db *gorm.DB) {
	customerRepo := NewCustomerRepository(db)
	customerService := NewCustomerService(db, customerRepo)
	customerHandler := NewCustomerHandler(customerService)

	CustomerRoute(app, customerHandler)
}
