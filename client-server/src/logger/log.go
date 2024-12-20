package logger

import (
	"go.uber.org/zap"
	"log"
)

const (
	DEBUG = iota
	INFO
)

// TODO: optional build
var Log *zap.Logger = NewLogger(DEBUG)

func NewLogger(level int) *zap.Logger {
	switch level {
	case DEBUG:
		loggerInstance, err := zap.NewDevelopment()
		if err != nil {
			log.Fatal(err)
		}
		return loggerInstance
	case INFO:
		loggerInstance, err := zap.NewProduction()
		if err != nil {
			log.Fatal(err)
		}
		return loggerInstance
	default:
		loggerInstance, err := zap.NewProduction()
		if err != nil {
			log.Fatal(err)
		}
		return loggerInstance
	}
}
