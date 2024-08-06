package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
	Staus string `json:"status" bson:"status"`
	Date primitive.DateTime `json:"date" bson:"date"`
	Creater primitive.ObjectID `json:"creator_id,omitempty" bson:"creator_id,omitempty"`
}
