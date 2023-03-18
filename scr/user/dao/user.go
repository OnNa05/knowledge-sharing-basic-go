package dao

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Name     string             `json:"name" bson:"name"`
	Email    string             `json:"email" bson:"email" validate:"required"`
	Password string             `json:"password" bson:"password" validate:"required"`
	Role     string             `json:"role" bson:"role"`
	CreateAt time.Time          `json:"create_at" bson:"create_at"`
}

type CreateUserResponse struct {
	ID string `json:"id"`
}

type UpdateUserResponse struct {
	ID string `json:"id"`
}

type DeleteUserRequest struct {
	ID primitive.ObjectID `json:"_id" validate:"required"`
}

type UpdateUserRequest struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id" validate:"required"`
	Name     string             `json:"name" bson:"name"`
	Email    string             `json:"email" bson:"email"`
	Password string             `json:"password" bson:"password"`
	Role     string             `json:"role" bson:"role"`
}
