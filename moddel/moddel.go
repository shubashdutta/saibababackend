package moddel

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	First_Name string             `json:"firstName" bson:"firstName"`
	Last_Name  string             `json:"lastName" bson:"lastName"`
	Email      string             `json:"email" bson:"email"`
	Phone      string             `json:"phone" bson:"phone"`
	Password   string             `json:"password" bson:"password"`
}
