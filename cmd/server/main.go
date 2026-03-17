package main

import (
	"log"

	"github.com/doitung/DoiTung-service/internal/config"
	"github.com/gofiber/fiber/v2"
	"github.com/doitung/DoiTung-service/internal/modules/auth"
)


func main() {

	app := fiber.New()

	config.ConnectDatabase()
	auth.RegisterRoutes(app)

	log.Fatal(app.Listen(":8080"))
}