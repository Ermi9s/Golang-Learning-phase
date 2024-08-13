package domain

import (
	"github.com/dgrijalva/jwt-go"
)

type UserClaims struct {
	jwt.StandardClaims
	ID    string `json:"_id"`
	Email string             `json:"email"`
	Is_admin bool            `json:"is_admin"`
}

