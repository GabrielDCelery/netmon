package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func NewLogger() (*zap.Logger, error) {
	env := os.Getenv("ENV")
	switch env {
	case "prod":
		config := zap.NewProductionConfig()
		return config.Build()
	default:
		w := zapcore.AddSync(&lumberjack.Logger{
			Filename:   "./logs/dev.logs",
			MaxSize:    1,
			MaxBackups: 1,
			MaxAge:     7,
			Compress:   false,
		})
		core := zapcore.NewCore(
			zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
			w,
			zap.DebugLevel,
		)
		return zap.New(core), nil
	}
}
