package routers

import (
	"github.com/gofiber/fiber/v2"
)

func EnrollRouter(app *fiber.App) {
	// api base path = /
	apiRouter := app.Group("/")
	initAuthRouter(apiRouter)
	initGroupRouter(apiRouter)
	initUserRouter(apiRouter)
	initCategoryRouter(apiRouter)
	initSubjectRouter(apiRouter)
	initStudylogRouter(apiRouter)
	initInviteRouter(apiRouter)
	initTimerRouter(apiRouter)
}
