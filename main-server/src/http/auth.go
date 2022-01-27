package http

import (
	"github.com/golang-jwt/jwt"
)

type AdminClaims struct {
	Username string `json:"username"`
	Office   int16  `json:"office"`
	jwt.StandardClaims
}

type UserClaims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}
