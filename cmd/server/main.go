package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
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

	// TODO add port config
	port := 8080
	logger.Fatal("Failed to start HTTP server", zap.Error(e.Start(fmt.Sprintf(":%d", port))))
}
