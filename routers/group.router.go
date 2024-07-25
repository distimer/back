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
	groupRouter.Delete("/:id", groupctrl.DeleteGroup)

	// group join
	groupRouter.Post("/join", groupctrl.JoinGroup)

	// group member
	groupRouter.Get("/member/:id", groupctrl.GetAllGroupMembers)

	// group policy
	groupRouter.Put("/policy/:id", groupctrl.ModifyGroupPolicy)

	// group invite
	groupRouter.Post("/invite/:id", groupctrl.InviteToGroup)
}
