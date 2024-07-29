package routers

import (
	"github.com/gofiber/fiber/v2"
	"pentag.kr/distimer/controllers/categoryctrl"
	"pentag.kr/distimer/middlewares"
)

func initCategoryRouter(router fiber.Router) {
	categoryRouter := router.Group("/category")

	categoryRouter.Use(middlewares.JWTMiddleware)

	// category
	categoryRouter.Get("/", categoryctrl.GetCategoryList)
	categoryRouter.Post("/", categoryctrl.CreateCategory)
	categoryRouter.Put("/:id", categoryctrl.ModifyCategory)
	categoryRouter.Delete("/:id", categoryctrl.DeleteCategory)
}
