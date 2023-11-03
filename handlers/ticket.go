package handler

import "github.com/gofiber/fiber/v2"

func GetAllTickets(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
func GetTicket(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
