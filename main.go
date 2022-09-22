package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initialize a Fiber application
	app := fiber.New()

	// Get function to route to / path and an handler function that returns a JSON
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"data": "Server has began",
		})
	})

	// Get function to route to / path and an handler function that returns a JSON

	app.Listen(":6000")
}
