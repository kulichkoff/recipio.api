package routes

import (
	"github.com/labstack/echo/v4"
	"recip.io/api/internal/handlers"
)

func activateAuthRoutes(e *echo.Echo) {
	h := handlers.NewAuthHandler()
	authGroup := e.Group("/auth")

	authGroup.POST("/sign-up", h.Register)
	authGroup.POST("/login", h.Login)
}
