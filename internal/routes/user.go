package routes

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"recip.io/api/internal/auth"
	"recip.io/api/internal/handlers"
)

func activateUserRoutes(e *echo.Echo) {
	h := handlers.NewUserHandler()
	rg := e.Group("/user")
	rg.Use(echojwt.WithConfig(auth.EchoJWTConfig()))

	rg.GET("/me", h.GetMe)
}
