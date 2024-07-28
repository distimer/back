package routers

import (
	"github.com/gofiber/fiber/v2"
	"pentag.kr/distimer/controllers/subjectctrl"
	"pentag.kr/distimer/middlewares"
)

func initSubjectRouter(router fiber.Router) {
	subjectRouter := router.Group("/subject")

	subjectRouter.Use(middlewares.JWTMiddleware)

	// subject
	subjectRouter.Post("/", subjectctrl.CreateSubject)
	subjectRouter.Delete("/:id", subjectctrl.DeleteSubject)
	subjectRouter.Put("/:id", subjectctrl.ModifySubjectInfo)

	subjectRouter.Get("/studylog/:id", subjectctrl.GetStudyLogWithSubject)
}
