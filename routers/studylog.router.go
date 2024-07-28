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

}
