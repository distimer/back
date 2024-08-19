package logger

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var MyLogger *zap.Logger

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
	config.Encoding = "console"
	logger, err := config.Build()
	if err != nil {
		panic(err)
	}
	MyLogger = logger
}

func Fatal(err error) {
	MyLogger.Fatal(err.Error())
}

func CtxError(c *fiber.Ctx, err error) {
	// METHOD IP URL ERROR
	MyLogger.Error("Internal Server Error", zap.String("METHOD", c.Method()), zap.String("IP", func() string {
		// cloudflare ip
		if ip := c.Get("CF-Connecting-IP"); ip != "" {
			return ip
		}
		return c.IP()
	}()), zap.String("URL", c.OriginalURL()), zap.Error(err))
}

func Error(err error) {
	MyLogger.Error(err.Error())
}

// unused but keep it for future use
// func Warn(msg string) {
// 	myLogger.Warn(msg)
// }

func Info(msg string) {
	MyLogger.Info(msg)
}

func Debug(msg string) {
	MyLogger.Debug(msg)
}
