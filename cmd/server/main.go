package main

import (
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
}
