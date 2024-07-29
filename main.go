package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"pentag.kr/distimer/configs"
	"pentag.kr/distimer/db"
	"pentag.kr/distimer/routers"
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

	// Load environment variables
	configs.LoadEnv()

	// Initialize logger
	logger.InitLogger(configs.Env.LogLevel)
	logger.Debug(time.Now().Format(time.RFC3339))
	// Connect database client
	db.ConnectDBClient()

	app := fiber.New()
	if configs.Env.LogLevel == "DEBUG" {
		swaggerConf := swagger.ConfigDefault
		swaggerConf.CustomStyle = configs.SwaggerDarkStyle
		app.Get("/swagger/*", swagger.New(swaggerConf)) // default
	}

	// Register routers
	routers.EnrollRouter(app)

	logger.Fatal(app.Listen(":3000"))
}
