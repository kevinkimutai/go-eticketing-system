package handler

import (
	"e-ticketing/db"
	"e-ticketing/model"

	"github.com/gofiber/fiber/v2"
)

func CreateEvent(c *fiber.Ctx) error {
	event := new(model.Event)

	if err := c.BodyParser(event); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	if event.Name == "" || event.Venue == "" || event.Location == "" || event.DateOfEvent == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
			"error":   "Missing Required Fields Of Event",
		})
	}

	userId := c.Locals("userId").(float64)
	convertedUserId := uint(userId)

	// 	db.Create(&Event{
	//   Name: event.Name,
	//   Venue:event.Venue,
	//   Location:event.Location,
	//   Price:event.Price,
	//   BannerUrl:event.BannerUrl,
	//   VendorId:userId,

	//   TicketType: []TicketType{event.TicketType}
	// })

	newEvent := &model.Event{
		Name:      event.Name,
		Venue:     event.Venue,
		Location:  event.Location,
		Price:     event.Price,
		BannerUrl: event.BannerUrl,
		VendorID:  convertedUserId,
	}

	newEvent.TicketType = append(newEvent.TicketType, event.TicketType...)

	//CREATE EVENT
	if err := db.DB.Create(newEvent).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(newEvent)

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
