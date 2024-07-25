package routers

import (
	"github.com/gofiber/fiber/v2"
)

func EnrollRouter(app *fiber.App) {
	// api base path = /api
	apiRouter := app.Group("/")
	initAuthRouter(apiRouter)
	initGroupRouter(apiRouter)
	initUserRouter(apiRouter)
}
