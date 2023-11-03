package handler

import "github.com/gofiber/fiber/v2"

func GetAllUsers(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func GetUser(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func UpdateUser(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func GetLoggedInUser(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func UpdateLoggedInUser(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func DeleteLoggedInUser(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
func GetAllTicketsByLoggedInUser(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
func GetAllTransactionsByLoggedInUser(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
