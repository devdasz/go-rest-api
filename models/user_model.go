package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// a struct with required properties. We added omitempty and validate:"required" to the struct tag to tell Fiber to ignore empty fields and make the field required, respectively.
type User struct {
	Id          primitive.ObjectID `json:"id,omitempty"`
	Name        string             `json:"name,omitempty" validate:"required"`
	Dob         string             `json:"dob,omitempty" validate:"required"`
	Address     string             `json:"address,omitempty" validate:"required"`
	Description string             `json:"description,omitempty" validate:"required"`
	CreatedAt   string             `json:"createdAt,omitempty" `
}
