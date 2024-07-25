package routers

import (
	"github.com/gofiber/fiber/v2"
	"pentag.kr/distimer/controllers/authctrl"
)

func initAuthRouter(router fiber.Router) {
	authRouter := router.Group("/auth")

	// refresh
	authRouter.Post("/refresh", authctrl.Refresh)
	// logout
	authRouter.Delete("/logout", authctrl.Logout)

	// oauth
	authRouter.Post("/oauth/google", authctrl.GoogleOauthLogin)
	authRouter.Post("/oauth/apple", authctrl.AppleOauthLogin)
}
