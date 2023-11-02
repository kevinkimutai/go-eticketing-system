package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	//Init Project
	app := fiber.New()
	//Logger middleware
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to ETICKETING")
	})

	app.Listen(":3000")
}
