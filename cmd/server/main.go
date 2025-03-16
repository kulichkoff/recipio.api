package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	emiddleware "github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
	"recip.io/api/internal/db"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	logger.Info("Logger initialised")

	if err := db.ConnectPostgres(logger); err != nil {
		panic(err)
	}
	logger.Info("DB connected")

	e := echo.New()
	e.HideBanner = true

	e.Use(emiddleware.RequestLoggerWithConfig(emiddleware.RequestLoggerConfig{
		LogURI:     true,
		LogLatency: true,
		LogStatus:  true,
		LogValuesFunc: func(c echo.Context, v emiddleware.RequestLoggerValues) error {
			logger.Info("request",
				zap.String("URI", v.URI),
				zap.Int("status", v.Status),
			)

			return nil
		},
	}))
	e.Use(emiddleware.Recover())
	e.Use(emiddleware.Gzip())
	e.Use(emiddleware.CORSWithConfig(emiddleware.DefaultCORSConfig))

	// TODO add port config
	port := 8080
	logger.Fatal("Failed to start HTTP server", zap.Error(e.Start(fmt.Sprintf(":%d", port))))
}
