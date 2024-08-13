package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Model interface {
	Valid() bool
}

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	UserName string             `json:"username,omitempty" bson:"username,omitempty"`
	Email    string             `json:"email,omitempty" bson:"email,omitempty"`
	Password string             `json:"password,omitempty" bson:"password,omitempty"`
	Is_admin bool               `json:"is_admin,omitempty" bson:"is_admin,omitempty"`
}

type AuthUser struct {
	ID       string 			`bson:"_id,omitempty" json:"_id,omitempty"`
	UserName string             `json:"username,omitempty" bson:"username,omitempty"`
	Email    string             `json:"email,omitempty" bson:"email,omitempty"`
	Password string             `json:"password,omitempty" bson:"password,omitempty"`
	Is_admin bool               `json:"is_admin,omitempty" bson:"is_admin,omitempty"`
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
