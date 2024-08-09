package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Model interface {
	Valid() bool
}

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	UserName string             `json:"username" bson:"username"`
	Email    string             `json:"email" bson:"email"`
	Password string             `json:"password" bson:"password"`
	Is_admin bool               `json:"is_admin" bson:"is_admin"`
}

type AuthUser struct {
	UserName string             `json:"username" bson:"username"`
	Email    string             `json:"email" bson:"email"`
	Password string             `json:"password" bson:"password"`
}

type LoggedInUser struct {
	ID       string 			`json:"id,omitempty" bson:"_id,omitempty"`
	UserName string             `json:"username" bson:"username"`
	Email    string             `json:"email" bson:"email"`
	Password string             `json:"password" bson:"password"`
	Is_admin bool               `json:"is_admin" bson:"is_admin"`
}

type Task struct {
	ID       	primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	Staus       string             `json:"status" bson:"status"`
	Date        primitive.DateTime `json:"date" bson:"date"`
	DueDate     primitive.DateTime `json:"duedate" bson:"duedate"`
	Creator     primitive.ObjectID `json:"creator_id,omitempty" bson:"creator_id,omitempty"`
}

