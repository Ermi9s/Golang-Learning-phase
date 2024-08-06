package models

import (
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserClaims struct {
	jwt.StandardClaims
	ID    primitive.ObjectID   `json:"_id"`
	Email string `json:"email"`
}

type User struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserName string             `json:"username" bson:"username"`
	Email    string             `json:"email" bson:"email"`
	Password string             `json:"password" bson:"password"`
	Is_admin bool               `json:"is_admin" bson:"is_admin"`
}

type LogIN struct {
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