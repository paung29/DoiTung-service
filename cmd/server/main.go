package main

import (
	"log"

	"github.com/doitung/DoiTung-service/internal/config"
	"github.com/doitung/DoiTung-service/internal/modules/account"
	"github.com/doitung/DoiTung-service/internal/modules/auth"
	"github.com/doitung/DoiTung-service/internal/modules/cluster"
	"github.com/doitung/DoiTung-service/internal/modules/customer"
	exportdata "github.com/doitung/DoiTung-service/internal/modules/exportdata"
	"github.com/doitung/DoiTung-service/internal/modules/forms/flower"
	harvestgrading "github.com/doitung/DoiTung-service/internal/modules/forms/harvestGrading"
	"github.com/doitung/DoiTung-service/internal/modules/forms/pod"
	"github.com/doitung/DoiTung-service/internal/modules/forms/pollination"
	preharvest "github.com/doitung/DoiTung-service/internal/modules/forms/preHarvest"
	"github.com/doitung/DoiTung-service/internal/modules/pole"
	"github.com/doitung/DoiTung-service/internal/modules/stock"
	"github.com/doitung/DoiTung-service/internal/modules/warehouse"
	"github.com/doitung/DoiTung-service/internal/modules/year"
	"github.com/doitung/DoiTung-service/internal/modules/zone"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found (this is okay in Docker)")
	}
}

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
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000", // frontend URL
		AllowCredentials: true,                    // ✅ REQUIRED
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
	}))

	auth.Setup(app, config.DB)
	account.Setup(app, config.DB)
	year.Setup(app, config.DB)
	zone.Setup(app, config.DB)
	pole.Setup(app, config.DB)
	cluster.Setup(app, config.DB)
	flower.Setup(app, config.DB)
	pollination.Setup(app, config.DB)
	pod.Setup(app, config.DB)
	preharvest.Setup(app, config.DB)
	harvestgrading.Setup(app, config.DB)
	warehouse.Setup(app, config.DB)
	customer.Setup(app, config.DB)
	stock.Setup(app, config.DB)
	exportdata.Setup(app, config.DB)

	log.Fatal(app.Listen(":8080"))
}
