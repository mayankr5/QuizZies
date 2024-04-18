package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/mayankr5/quizzies/app/router"
	"github.com/mayankr5/quizzies/database"
)

func Init() {
	app := fiber.New()
	app.Use(logger.New())

	router.SetupRoutes(app)
	database.Connect()
	app.Listen(":3000")
}
