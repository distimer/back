package routers

import (
	"github.com/gofiber/fiber/v2"
	"pentag.kr/distimer/controllers/pushctrl"
)

func initPushRouter(router fiber.Router) {
	pushRouter := router.Group("/push")

	pushRouter.Post("/start_token", pushctrl.UpsertStartToken)
	pushRouter.Post("/update_token", pushctrl.UpserUpdateToken)

}
