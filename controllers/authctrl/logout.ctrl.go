package authctrl

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"pentag.kr/distimer/db"
	"pentag.kr/distimer/ent"
	"pentag.kr/distimer/utils/dto"
	"pentag.kr/distimer/utils/logger"
)

func Logout(c *fiber.Ctx) error {
	data := new(refreshTokenDTO)
	if err := dto.Bind(c, data); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err,
		})
	}
	dbConn := db.GetDBClient()
	refreshToken := uuid.MustParse(data.RefreshToken)

	err := dbConn.RefreshToken.DeleteOneID(refreshToken).Exec(context.Background())
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
