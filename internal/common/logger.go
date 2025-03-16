package common

import "go.uber.org/zap"

var logger *zap.Logger

func SetLogger(zapLogger *zap.Logger) {
	logger = zapLogger
}

func GetLogger() *zap.Logger {
	return logger
}
