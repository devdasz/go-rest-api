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
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// create a collection
var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")

// validate models
var validate = validator.New()

// a CreateUser function that returns error if failed to operate
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

// a GetUser function that returns error if failed to operate
func GetAUser(c *fiber.Ctx) error {
	// defined a timeout of 10 seconds when inserting user into the document
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	userId := c.Params("userId")
	var user models.User
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(userId)

	// search for the user in database
	err := userCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&user)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}
	return c.Status(http.StatusOK).JSON(responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": user}})
}

// a EditUser function that returns error if failed to operate

func EditAUser(c *fiber.Ctx) error {
	// defined a timeout of 10 seconds when inserting user into the document
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	userId := c.Params("userId")
	var user models.User
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(userId)

	// validate the requets body
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}
	//use the validator library to validate required fields
	if validationErr := validate.Struct(&user); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}
	// set fields value for updation
	update := bson.M{
		"name":        user.Name,
		"dob":         user.Dob,
		"address":     user.Address,
		"description": user.Description,
	}

	// try to update on database
	result, err := userCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})

	// show error if failed
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//get updated user details
	var updatedUser models.User
	if result.MatchedCount == 1 {
		err := userCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updatedUser)

		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}
	}
	return c.Status(http.StatusOK).JSON(responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": updatedUser}})
}

// a DeleteAUser function that returns error if failed to operate
func DeleteAUser(c *fiber.Ctx) error {
	// defined a timeout of 10 seconds when inserting user into the document
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	userId := c.Params("userId")
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(userId)
	// try to delete on database
	result, err := userCollection.DeleteOne(ctx, bson.M{"id": objId})

	// if failed show error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	// check if user vaild
	if result.DeletedCount < 1 {
		return c.Status(http.StatusNotFound).JSON(
			responses.UserResponse{Status: http.StatusNotFound, Message: "error", Data: &fiber.Map{"data": "User with specified ID not found!"}},
		)
	}

	// show deletion success message
	return c.Status(http.StatusOK).JSON(
		responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": "User successfully deleted!"}},
	)
}

func GetAllUsers(c *fiber.Ctx) error {
	// defined a timeout of 10 seconds when inserting user into the document
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	// a slice of user type
	var users []models.User
	defer cancel()

	// fetch all users - the list of users
	results, err := userCollection.Find(ctx, bson.M{})

	// show error if failed
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}
	//reading from the db in an optimal way
	defer results.Close(ctx)
	// read the retuned list optimally using the Next attribute method to loop through the returned list of users
	for results.Next(ctx) {
		var singleUser models.User
		if err = results.Decode(&singleUser); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}

		users = append(users, singleUser)
	}

	return c.Status(http.StatusOK).JSON(
		responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": users}},
	)
}
