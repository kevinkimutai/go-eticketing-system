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
	app.Patch("/user/:id")
	// app.Delete("/user/:id", DeleteUser)
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
