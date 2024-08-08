package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type AuthUser struct {
	UserName string             `json:"username" bson:"username"`
	Email    string             `json:"email" bson:"email"`
	Password string             `json:"password" bson:"password"`
}

type LoggedInUser struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserName string             `json:"username" bson:"username"`
	Email    string             `json:"email" bson:"email"`
	Is_admin bool               `json:"is_admin" bson:"is_admin"`
}


