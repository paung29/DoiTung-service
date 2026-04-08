package main

import (
	"log"

	"github.com/doitung/DoiTung-service/internal/config"
	"github.com/doitung/DoiTung-service/internal/modules/account"
	"github.com/doitung/DoiTung-service/internal/modules/auth"
	"github.com/doitung/DoiTung-service/internal/modules/cluster"
	"github.com/doitung/DoiTung-service/internal/modules/forms/flower"
	"github.com/doitung/DoiTung-service/internal/modules/year"
	"github.com/doitung/DoiTung-service/internal/modules/zone"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}

			return c.Status(code).JSON(fiber.Map{
				"message": err.Error(),
			})
		},
	})

	config.ConnectDatabase()

	app.Use(logger.New())

	auth.Setup(app, config.DB)
	account.Setup(app, config.DB)
	year.Setup(app, config.DB)
	zone.Setup(app, config.DB)
	cluster.Setup(app, config.DB)
	flower.Setup(app, config.DB)

	log.Fatal(app.Listen(":8080"))
}
