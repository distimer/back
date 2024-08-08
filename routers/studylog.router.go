package routers

import (
	"github.com/gofiber/fiber/v2"
	"pentag.kr/distimer/controllers/studylogctrl"
	"pentag.kr/distimer/middlewares"
)

func initStudylogRouter(router fiber.Router) {
	studylogRouter := router.Group("/studylog")

	studylogRouter.Use(middlewares.JWTMiddleware)

	// studylog
	studylogRouter.Get("/", studylogctrl.GetAllMyStudyLogs)
	studylogRouter.Post("/", studylogctrl.CreateStudyLog)
	studylogRouter.Put("/:id", studylogctrl.ModifyStudyLog)
	studylogRouter.Delete("/:id", studylogctrl.DeleteStudyLog)

	// studylog statistics
	studylogRouter.Get("/statistics/date", studylogctrl.GetStatisticsWithDate)
	studylogRouter.Get("/statistics/term", studylogctrl.GetStatisticsWithTerm)

	// studylog with group
	studylogRouter.Get("/group/statistics/date/:id", studylogctrl.GroupMemberStatisticsByDate)

	// studylog with subject
	studylogRouter.Get("/subject/:id", studylogctrl.GetStudyLogWithSubject)
}
