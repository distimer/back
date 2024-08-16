package routers

import (
	"github.com/gofiber/fiber/v2"
	"pentag.kr/distimer/controllers/subjectctrl"
	"pentag.kr/distimer/middlewares"
)

func initSubjectRouter(router fiber.Router) {
	subjectRouter := router.Group("/subject")

	subjectRouter.Use(middlewares.JWTMiddleware)

	// create batch subject
	subjectRouter.Post("/batch", subjectctrl.CreateBatchSubject)

	// subject
	subjectRouter.Post("/:id", subjectctrl.CreateSubject)
	subjectRouter.Delete("/:id", subjectctrl.DeleteSubject)
	subjectRouter.Put("/:id", subjectctrl.ModifySubjectInfo)

	// subject order
	subjectRouter.Patch("/order", subjectctrl.SubjectOrderModify)
}
