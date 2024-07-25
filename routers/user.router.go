package routers

import (
	"github.com/gofiber/fiber/v2"
	"pentag.kr/distimer/controllers/userctrl"
	"pentag.kr/distimer/middlewares"
)

func initUserRouter(router fiber.Router) {
	userRouter := router.Group("/user")

	userRouter.Use(middlewares.JWTMiddleware)

	userRouter.Get("/", userctrl.GetMyUserInfo)
	userRouter.Put("/", userctrl.ModifyUserInfo)
}
