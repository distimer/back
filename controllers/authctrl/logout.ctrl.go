package authctrl

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"pentag.kr/distimer/db"
	"pentag.kr/distimer/ent"
	"pentag.kr/distimer/utils/logger"
)

func Logout(c *fiber.Ctx) error {
	refreshTokenStr := c.Query("refresh_token")
	refreshToken, err := uuid.Parse(refreshTokenStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid refresh token",
		})
	}
	dbConn := db.GetDBClient()

	err = dbConn.RefreshToken.DeleteOneID(refreshToken).Exec(context.Background())
	if err != nil {
		if ent.IsNotFound(err) {
			return c.Status(401).JSON(fiber.Map{
				"error": "Invalid refresh token",
			})
		}
		logger.CtxError(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	return c.SendStatus(204)
}
