package main

import (
	"log"

	"github.com/doitung/DoiTung-service/internal/config"
	"github.com/gofiber/fiber/v2"
)


func main() {

	app := fiber.New()

	config.ConnectDatabase()

	log.Fatal(app.Listen(":8080"))
}