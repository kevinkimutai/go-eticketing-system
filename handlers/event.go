package handler

import (
	"e-ticketing/db"
	"e-ticketing/model"
	"time"

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

	newEvent := &model.Event{
		Name:        event.Name,
		Venue:       event.Venue,
		Location:    event.Location,
		DateOfEvent: event.DateOfEvent,
		Description: event.Description,
		Price:       event.Price,
		BannerUrl:   event.BannerUrl,
		VendorID:    convertedUserId,
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

func GetAllEvents(c *fiber.Ctx) error {
	req := c.Queries()
	//TODO ADD MORE QUERIES FOR DASHBOARD.
	var events []model.Event

	query := db.DB

	if req["dateOfEvent"] != "" {
		query = query.Where("date_of_event = ?", req["dateOfEvent"])
	}
	if req["search"] != "" {
		query = query.Where("name LIKE ? OR description LIKE ?", "%"+req["search"]+"%", "%"+req["search"]+"%")
	}
	if req["location"] != "" {
		query = query.Where("location = ?", req["location"])
	}
	query = query.Where("date_of_event > ?", time.Now())
	//TODO QUERY NEAREST EVENTS

	if err := query.Find(&events).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(events)
}

func GetEvent(c *fiber.Ctx) error {
	var event model.Event
	eventId := c.Params("eventId")

	if err := db.DB.First(&event, eventId).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(event)

}

func UpdateEvent(c *fiber.Ctx) error {
	userId := c.Locals("userId").(float64)
	convertedUserId := uint(userId)
	eventId := c.Params("eventId")

	event := new(model.Event)

	if err := db.DB.First(&event, eventId).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
	}

	if event.VendorID != convertedUserId {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorised",
			"error":   "Unauthorised,Only The Vendor Of The Event Can Update",
		})
	}

	//BodyParser
	if err := c.BodyParser(event); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	// Save the updated event back to the database
	if err := db.DB.Save(&event).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Event updated successfully",
		"event":   event,
	})

}

func DeleteEvent(c *fiber.Ctx) error {
	userId := c.Locals("userId").(float64)
	convertedUserId := uint(userId)
	eventId := c.Params("eventId")

	event := new(model.Event)

	if err := db.DB.First(&event, eventId).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
	}

	if event.VendorID != convertedUserId {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorised",
			"error":   "Unauthorised,Only The Vendor Of The Event Can Update",
		})
	}

	// Save the updated event back to the database
	if err := db.DB.Delete(&event).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusNoContent).Send(nil)
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
