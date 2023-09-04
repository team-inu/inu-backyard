package logger

import (
	"go.uber.org/zap"
)

func NewZapLogger() *zap.Logger {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	return logger
}
