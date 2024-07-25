package logger

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var myLogger *zap.Logger

func InitLogger(logLevel string) {
	var level zapcore.Level
	switch logLevel {
	case "DEBUG":
		level = zapcore.DebugLevel
	case "INFO":
		level = zapcore.InfoLevel
	case "WARN":
		level = zapcore.WarnLevel
	case "ERROR":
		level = zapcore.ErrorLevel
	case "FATAL":
		level = zapcore.FatalLevel
	default:
		level = zapcore.InfoLevel
	}

	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(level)
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	logger, err := config.Build()
	if err != nil {
		panic(err)
	}
	myLogger = logger
}

func Fatal(err error) {
	myLogger.Fatal(err.Error())
}

func Error(c *fiber.Ctx, err error) {
	// METHOD IP URL ERROR
	myLogger.Error("Internal Server Error", zap.String("METHOD", c.Method()), zap.String("IP", c.IP()), zap.String("URL", c.OriginalURL()), zap.Error(err))
}

// unused but keep it for future use
// func Warn(msg string) {
// 	myLogger.Warn(msg)
// }

func Info(msg string) {
	myLogger.Info(msg)
}

func Debug(msg string) {
	myLogger.Debug(msg)
}
