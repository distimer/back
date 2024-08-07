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
	studylogRouter.Delete("/:id", studylogctrl.DeleteStudyLog)

	// studylog statistics
	studylogRouter.Get("/statistics/date", studylogctrl.GetStatisticsWithDate)
	studylogRouter.Get("/statistics/term", studylogctrl.GetStatisticsWithTerm)

	studylogRouter.Get("/subject/:id", studylogctrl.GetStudyLogWithSubject)
}
