package routes

import "github.com/labstack/echo/v4"

func ActivateRoutes(e *echo.Echo) {
	activateAuthRoutes(e)
	activateUserRoutes(e)
}
