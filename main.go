package main

import (
	"time"

	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"pentag.kr/distimer/configs"
	"pentag.kr/distimer/db"
	"pentag.kr/distimer/routers"
	"pentag.kr/distimer/schedulers"
	"pentag.kr/distimer/utils/logger"

	_ "pentag.kr/distimer/docs"
)

// @title Distimer Swagger API
// @version	1.0
// @host localhost:3000
// @BasePath  /
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
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

	app := fiber.New()

	prometheus := fiberprometheus.New("my-service-name")
	prometheus.RegisterAt(app, "/metrics")
	app.Use(prometheus.Middleware)

	if configs.Env.LogLevel == "DEBUG" {
		swaggerConf := swagger.ConfigDefault
		swaggerConf.CustomStyle = configs.SwaggerDarkStyle
		app.Get("/swagger/*", swagger.New(swaggerConf)) // default
	}

	// Register routers
	routers.EnrollRouter(app)

	logger.Fatal(app.Listen(":3000"))
	schedulerObj.Stop()
}
