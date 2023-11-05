package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"e-ticketing/db"
	handler "e-ticketing/handlers"
)

func routers(app *fiber.App) {

	//Auth Routes
	app.Post("/auth/login", handler.Login)
	app.Post("/auth/signup", handler.SignUp)

	//User Routes
	app.Get("/users", handler.Protected, handler.Restricted("admin"), handler.GetAllUsers)
	app.Get("/user/:userId", handler.Protected, handler.Restricted("admin"), handler.GetUser)
	app.Patch("/user/:userId", handler.Protected, handler.Restricted("admin"), handler.UpdateUser)

	//Me Route
	app.Get("/me", handler.Protected, handler.GetLoggedInUser)
	app.Patch("/me", handler.Protected, handler.UpdateLoggedInUser)
	app.Delete("/me", handler.Protected, handler.DeleteLoggedInUser)
	app.Get("/me/tickets", handler.Protected, handler.GetAllTicketsByLoggedInUser)
	app.Get("/me/transactions", handler.Protected, handler.GetAllTransactionsByLoggedInUser)

	//Event Routes
	//Only user who created event can delete/update
	app.Post("/event", handler.Protected, handler.Restricted("user"), handler.CreateEvent)
	app.Get("/event", handler.GetAllEvents)
	app.Get("/event/:eventId", handler.Protected, handler.GetEvent)
	app.Patch("/event/:eventId", handler.Protected, handler.Restricted("user"), handler.UpdateEvent)
	app.Delete("/event/:eventId", handler.Protected, handler.Restricted("user"), handler.DeleteEvent)

	//Event-Ticket Routes
	app.Post("/event/:eventId/ticket", handler.Protected, handler.Restricted("user"), handler.BookEvent)
	app.Get("/event/:eventId/ticket", handler.Protected, handler.Restricted("user"), handler.GetAllTicketsByEvent)
	app.Post("/event/:eventId/ticket/:ticketId", handler.Protected, handler.Restricted("user"), handler.GetTicketByEvent)

	//Event-Ticket_Type Routes
	//Only Vendor oF Event can create  Ticket-Type
	app.Post("/event/:eventId/ticket-type", handler.Protected, handler.Restricted("user"), handler.CreateTicketType)
	app.Get("/event/:eventId/ticket-type", handler.Protected, handler.Restricted("user"), handler.GetAllTicketTypesOfEvent)
	app.Get("/event/:eventId/ticket-type/:ticketTypeId", handler.Protected, handler.Restricted("user"), handler.GetTicketTypeByEvent)
	app.Patch("/event/:eventId/ticket-type/:ticketTypeId", handler.Protected, handler.Restricted("user"), handler.UpdateTicketType)

	//Ticket Routes
	app.Get("/tickets", handler.Protected, handler.Restricted("admin"), handler.GetAllTickets)
	app.Get("/ticket/:ticketId", handler.Protected, handler.Restricted("admin"), handler.GetTicket)

	//Ticket_Type Routes
	app.Get("/ticketType", handler.Protected, handler.Restricted("admin"), handler.GetAllTicketTypes)
	app.Get("/ticketType/:ticketTypeId", handler.Protected, handler.Restricted("admin"), handler.GetTicketType)

}

func main() {
	//DB INIT
	db.InitMigration()

	//Init Project
	app := fiber.New()
	//Logger middleware
	app.Use(logger.New())

	//Routes
	routers(app)

	app.Listen(":3000")
}
