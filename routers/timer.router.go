package routers

import (
	"github.com/gofiber/fiber/v2"
	"pentag.kr/distimer/controllers/timerctrl"
	"pentag.kr/distimer/middlewares"
)

func initTimerRouter(router fiber.Router) {
	timerRouter := router.Group("/timer")

	timerRouter.Use(middlewares.JWTMiddleware)

	// timer
	timerRouter.Get("/", timerctrl.GetMyTimerInfo)
	timerRouter.Post("/", timerctrl.CreateTimer)
	timerRouter.Put("/", timerctrl.ModifyTimer)
	timerRouter.Delete("/", timerctrl.DeleteTimer)
}
