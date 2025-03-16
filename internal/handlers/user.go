package handlers

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"recip.io/api/internal/auth"
)

type userHandler struct{}

func NewUserHandler() *userHandler {
	return &userHandler{}
}

func (h *userHandler) GetMe(c echo.Context) error {
	userToken := c.Get("user").(*jwt.Token)
	userClaims := userToken.Claims.(*auth.JWTUserClaims)
	name := userClaims.Name
	return c.String(200, "Hello "+name)
}
