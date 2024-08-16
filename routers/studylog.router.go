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

	// studylog detail
	studylogRouter.Get("/detail/:id", studylogctrl.GetDetailByID)

	// studylog with date
	studylogRouter.Get("/date", studylogctrl.GetByDate)
	// studylog with term
	studylogRouter.Get("/term", studylogctrl.GetByTerm)

	// studylog statistics
	studylogRouter.Get("/statistics/date", studylogctrl.GetStatisticsByDate)
	studylogRouter.Get("/statistics/term", studylogctrl.GetStatisticsByTerm)

	// studylog with subject
	studylogRouter.Get("/subject/:id", studylogctrl.GetStudyLogWithSubject)

	// studylog with group
	studylogRouter.Get("/group/term/:group_id/:member_id", studylogctrl.GetByTermWithGroup)

	// studylog statistics with group
	studylogRouter.Get("/group/statistics/term/:group_id/:member_id", studylogctrl.GetStatisticsByTermWithGroup)
}
