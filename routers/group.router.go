package routers

import (
	"github.com/gofiber/fiber/v2"
	"pentag.kr/distimer/controllers/groupctrl"
	"pentag.kr/distimer/middlewares"
)

func initGroupRouter(router fiber.Router) {
	groupRouter := router.Group("/group")

	groupRouter.Use(middlewares.JWTMiddleware)

	// group
	groupRouter.Get("/", groupctrl.GetJoinedGroups)
	groupRouter.Post("/", groupctrl.CreateGroup)
	groupRouter.Put("/:id", groupctrl.ModifyGroupInfo)
	groupRouter.Delete("/:id", groupctrl.DeleteGroup)

	// group join
	groupRouter.Post("/join", groupctrl.JoinGroup)

	// group quit
	groupRouter.Delete("/quit/:id", groupctrl.QuitGroup)

	// group member
	groupRouter.Get("/member/:id", groupctrl.GetAllGroupMembers)

	// group member modify
	groupRouter.Put("/member/:group_id/:member_id", groupctrl.ModifyMember)

	// group nickname
	groupRouter.Patch("/nickname/:id", groupctrl.ModifyNickname)
}
