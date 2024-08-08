package domain

import (
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserClaims struct {
	jwt.StandardClaims
	ID    primitive.ObjectID `json:"_id"`
	Email string             `json:"email"`
	Is_admin bool            `json:"is_admin"`
}

