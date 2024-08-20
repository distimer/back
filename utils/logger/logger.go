package logger

import (
	"log/slog"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/grafana/loki-client-go/loki"
	slogloki "github.com/samber/slog-loki/v3"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"pentag.kr/distimer/configs"
)

var MyLogger *zap.Logger
var LokiLogger *slog.Logger

func InitLokiLogger() {
	// setup loki client
	config, _ := loki.NewDefaultConfig(configs.Env.LokiURL)
	config.TenantID = "distimer"
	client, _ := loki.New(config)

	logger := slog.New(slogloki.Option{Level: slog.LevelInfo, Client: client}.NewLokiHandler())
	logger = logger.
		With("app", "distimer").
		With("branch", configs.Env.Branch)
	LokiLogger = logger
}

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

	rotateLogger := &lumberjack.Logger{
		Filename:   "./log/log.log", // Or any other path
		MaxSize:    100,             // MB; after this size, a new log file is created
		MaxBackups: 30,              // Number of backups to keep
		MaxAge:     90,              // Days
		Compress:   true,            // Compress the backups using gzip
	}

	consoleSyncer := zapcore.AddSync(os.Stdout)
	fileSyncer := zapcore.AddSync(rotateLogger)

	consoleCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
		consoleSyncer,
		level,
	)
	fileCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
		fileSyncer,
		level,
	)
	core := zapcore.NewTee(fileCore, consoleCore)

	loggerZap := zap.New(core)
	defer loggerZap.Sync()
	MyLogger = loggerZap

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
