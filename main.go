package main

import (
	"time"

	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2"
	"pentag.kr/distimer/configs"
	"pentag.kr/distimer/db"
	"pentag.kr/distimer/middlewares"
	"pentag.kr/distimer/routers"
	"pentag.kr/distimer/schedulers"
	"pentag.kr/distimer/utils/logger"
)

func main() {

	location, err := time.LoadLocation("Asia/Seoul")
	if err != nil {
		panic(err)
	}

	// Set the timezone for the current process
	time.Local = location

	// Load environment variables
	configs.LoadEnv()

	// Initialize logger
	logger.InitLogger(configs.Env.LogLevel)
	logger.Debug(time.Now().Format(time.RFC3339))
	// Connect database client
	db.ConnectDBClient()

	schedulerObj := schedulers.GenerateSchedularObj()
	schedulerObj.Start()

	app := fiber.New(fiber.Config{
		ProxyHeader: "CF-Connecting-IP",
	})

	if configs.Env.Branch != "local" {
		logger.InitLokiLogger()
		app.Use(middlewares.LokiLoggerMiddleware)
	}

	app.Use(fiberzap.New(fiberzap.Config{
		Logger: logger.MyLogger,
	}))

	// Register routers
	routers.EnrollRouter(app)

	logger.Fatal(app.Listen(":3000"))
	schedulerObj.Stop()
}
