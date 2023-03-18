package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserEntitty struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Name     string             `json:"name" bson:"name"`
	Email    string             `json:"email" bson:"email"`
	Password string             `json:"password" bson:"password"`
	Role     string             `json:"role" bson:"role"`
	CreateAt time.Time          `json:"create_at" bson:"create_at"`
}
