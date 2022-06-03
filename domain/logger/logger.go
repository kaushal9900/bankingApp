package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	var err error
	config := zap.NewProductionConfig()
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig = encoderConfig
	log, err = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
}

func Info(message string, feilds ...zap.Field) {
	log.Info(message, feilds...)
}

func Debug(message string, feilds ...zap.Field) {
	log.Debug(message, feilds...)
}

func Error(message string, feilds ...zap.Field) {
	log.Error(message, feilds...)
}
