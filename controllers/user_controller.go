package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/devdasz/go-rest-api/configs"
	"github.com/devdasz/go-rest-api/models"
	"github.com/devdasz/go-rest-api/responses"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// create a collection
var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")

// validate models
var validate = validator.New()

// a CreateUser function that returns if failed to operate
func CreateUser(c *fiber.Ctx) error {

	// defined a timeout of 10 seconds when inserting user into the document
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	// user variable
	var user models.User
	defer cancel()

	// check request body
	if err := c.BodyParser(&user); err != nil {
		// if failed return error message
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})

	}
	// check required fileds
	if validationErr := validate.Struct(&user); validationErr != nil {
		// if failed return error message
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	newUser := models.User{
		Id:          primitive.NewObjectID(),
		Name:        user.Name,
		Dob:         user.Dob,
		Address:     user.Address,
		Description: user.Description,
		CreatedAt:   time.Now().UTC().GoString(),
	}
	// insert it into database , return error if failed
	result, err := userCollection.InsertOne(ctx, newUser)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}
	return c.Status(http.StatusCreated).JSON(responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": result}})

}
