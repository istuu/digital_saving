package structs

import (
	"time"
	"gopkg.in/mgo.v2/bson"
)

type Customer struct {
	ID bson.ObjectId `json:"id" bson:"_id"`
	CitizenId int `json:"citizen_id" bson:"citizen_id"`
	FirstName string `json:"first_name" bson:"first_name"`
	MidleName string `json:"midle_name" bson:"midle_name"`
	LastName string `json:"last_name" bson:"last_name"`
	Email string `json:"email" bson:"email"`
	ProductCode string `json:"product_code" bson:"product_code"`
	CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at"`
}