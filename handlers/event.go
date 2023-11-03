package handler

import "github.com/gofiber/fiber/v2"

func CreateEvent(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func GetEvent(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func UpdateEvent(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func DeleteEvent(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func BookEvent(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func GetAllTicketsByEvent(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func CreateTicketType(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func GetAllTicketTypesOfEvent(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func GetTicketTypeByEvent(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func UpdateTicketType(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func GetTicketByEvent(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
