package responses

import "github.com/gofiber/fiber/v2"

// a  struct with Status, Message, and Data property to represent the API response type
type UserResponse struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Data    *fiber.Map `json:"data"`
}
