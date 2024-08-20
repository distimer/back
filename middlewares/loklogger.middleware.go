package middlewares

import (
	"log/slog"
	"time"

	"github.com/gofiber/fiber/v2"
	"pentag.kr/distimer/utils/logger"
)

// lokkiloggger
// 1. TIME
// 2. METHOD
// 3. PATH
// 4. STATUS
// 5. LATENCY
// 6. IP

func LokiLoggerMiddleware(c *fiber.Ctx) error {

	lokiLogger := logger.LokiLogger

	now := c.Context().Time()
	method := c.Method()
	path := c.Path()
	ip := c.IP()

	err := c.Next()

	status := c.Response().StatusCode()
	latency := time.Since(now).Microseconds()

	requestAttributes := []slog.Attr{
		slog.Time("time", now),
		slog.String("method", method),
		slog.String("path", path),
		slog.String("ip", ip),
		slog.Int("status", status),
		slog.Int64("latency", latency),
	}
	logErr := err
	if logErr == nil {
		logErr = fiber.NewError(status)
	}
	level := slog.LevelInfo
	msg := "Incoming request"
	if status >= 500 {
		level = slog.LevelError
		msg = logErr.Error()
	}
	lokiLogger.LogAttrs(c.UserContext(), level, msg, requestAttributes...)

	return err
}
