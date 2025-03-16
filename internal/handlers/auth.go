package handlers

import (
	"errors"

	"github.com/labstack/echo/v4"
	"recip.io/api/internal/auth"
	"recip.io/api/internal/models"
)

type authHandler struct {
	authService auth.AuthService
}

func NewAuthHandler() *authHandler {
	return &authHandler{
		authService: auth.NewService(),
	}
}

func (h *authHandler) Register(c echo.Context) error {
	return c.String(200, "OK")
}

func (h *authHandler) Login(c echo.Context) error {
	dto := &models.LoginDTO{}
	if err := c.Bind(dto); err != nil {
		return echo.ErrBadRequest
	}
	token, err := h.authService.Authorize(dto.Email, dto.Password)
	if err != nil {
		if errors.Is(err, auth.ErrUnauthorized) {
			return echo.ErrUnauthorized
		}
		return echo.ErrInternalServerError
	}

	return c.JSON(200, echo.Map{
		"authToken": token,
	})
}
