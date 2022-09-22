package routes

import (
	"github.com/devdasz/go-rest-api/controllers"
	"github.com/gofiber/fiber/v2"
)

// create user routes
func UserRoute(app *fiber.App) {
	//All routes related to users comes here

	// create a user route
	app.Post("/user", controllers.CreateUser)
	// get a user route
	app.Get("/user/:userId", controllers.GetAUser)
	// update a user route
	app.Put("/user/:userId", controllers.EditAUser)
}
