package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct { //to define our own datatype
	ID            primitive.ObjectID/*used for unique identifiers */ `bson:"id"` //struct tag
	First_name    *string                                                        `json:"first_name" validate:"required, min=2, max=100"`
	Last_name     *string                                                        `json:"last_name" validate:"required, min=2, max=100"`
	Password      *string                                                        `json:"Password" validate:"required, min=6"`
	Email         *string                                                        `json:"email" validate:"email, required"` //email is a validation type
	Phone         *string                                                        `json:"phone" validate:"required, max=10"`
	Token         *string                                                        `json:"token"`
	User_Type     *string                                                        `json:"user_type" validate:"required, eq=ADMIN|eq=USER"`
	Refresh_Token *string                                                        `json:"refresh_token"`
	Created_At    time.Time                                                      `json:"created_at"`
	Updated_At    time.Time                                                      `json:"updated_at"`
	User_id       string                                                         `json:"user_id"`
}
