package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserEntitty struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Email       string             `json:"email" bson:"email"`
	Password    string             `json:"password" bson:"password"`
	Role        string             `json:"role" bson:"role"`
	ActiveToken string             `json:"active_token" bson:"active_token"`
	AccessToken string             `json:"access_token" bson:"access_token"`
	CreateDate  time.Time          `json:"create_date" bson:"create_date"`
}
