package handler

import "github.com/gofiber/fiber/v2"

func GetAllTicketTypes(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func GetTicketType(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
