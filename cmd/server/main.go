package main

import (
	"log"

	"github.com/doitung/DoiTung-service/internal/config"
	"github.com/doitung/DoiTung-service/internal/modules/account"
	"github.com/doitung/DoiTung-service/internal/modules/auth"
	"github.com/doitung/DoiTung-service/internal/modules/year"
	"github.com/gofiber/fiber/v2"
)


func main() {

	app := fiber.New()

	config.ConnectDatabase()
	
	auth.Setup(app, config.DB)
	account.Setup(app, config.DB)
	year.Setup(app, config.DB)

	log.Fatal(app.Listen(":8080"))
}