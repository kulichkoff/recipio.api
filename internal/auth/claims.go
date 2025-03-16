package auth

import "github.com/golang-jwt/jwt/v5"

type JWTUserClaims struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}
