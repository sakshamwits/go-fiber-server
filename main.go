package main

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-server/database"
	"go-fiber-server/router"
)

func main() {
	app := fiber.New()

	database.ConnectDB()

	router.SetupRoutes(app)

	app.Listen(":3000")
}
