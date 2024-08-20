package middlewares

import (
	"fmt"
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

const (
	lokiLoggerFormat = `{"time":"%s","method":"%s","path":"%s","status":"%d","latency":%dus,"ip":"%s"}`
)

func LokiLoggerMiddleware(c *fiber.Ctx) error {

	lokiLogger := logger.LokiLogger

	now := time.Now().Format(time.RFC3339)
	method := c.Method()
	path := c.Path()
	// get cf-connecting-ip
	ip := c.IP()
	err := c.Next()
	status := c.Response().StatusCode()
	latency := time.Since(c.Context().Time()).Microseconds()

	lokiLogger.Info(fmt.Sprintf(lokiLoggerFormat, now, method, path, status, latency, ip))

	return err
}
