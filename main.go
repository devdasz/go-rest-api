package main

import (
	"github.com/devdasz/go-rest-api/configs"
	"github.com/devdasz/go-rest-api/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initialize a Fiber application
	app := fiber.New()

	//run database
	configs.ConnectDB()

	//routes
	routes.UserRoute(app)

	app.Listen(":6000")
}
