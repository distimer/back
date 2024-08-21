package timerctrl

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"pentag.kr/distimer/db"
	"pentag.kr/distimer/ent/timer"
	"pentag.kr/distimer/ent/user"
	"pentag.kr/distimer/middlewares"
	"pentag.kr/distimer/utils/logger"
)

func DeleteTimer(c *fiber.Ctx) error {

	userID := middlewares.GetUserIDFromMiddleware(c)

	dbConn := db.GetDBClient()

	count, err := dbConn.Timer.Delete().Where(timer.HasUserWith(user.ID(userID))).Exec(context.Background())
	if err != nil {
		logger.CtxError(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	} else if count == 0 {
		return c.Status(404).JSON(fiber.Map{
			"error": "Timer not found",
		})
	}
	return c.SendStatus(204)
}
