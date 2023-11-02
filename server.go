package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"e-ticketing/db"
)

func main() {
	//DB INIT
	db.InitMigration()

	//Init Project
	app := fiber.New()
	//Logger middleware
	app.Use(logger.New())

	//Routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}
