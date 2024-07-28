package routers

import (
	"github.com/gofiber/fiber/v2"
	"pentag.kr/distimer/controllers/invitectrl"
	"pentag.kr/distimer/middlewares"
)

func initInviteRouter(router fiber.Router) {
	inviteRouter := router.Group("/invite")

	inviteRouter.Use(middlewares.JWTMiddleware)

	// group invite
	inviteRouter.Get("/:code", invitectrl.GetInviteCodeInfo)
	inviteRouter.Get("/group/:id", invitectrl.GetInviteCodeList)
	inviteRouter.Post("/group/:id", invitectrl.InviteToGroup)
	inviteRouter.Delete("/group/:id/:code", invitectrl.DeleteInviteCode)
}
