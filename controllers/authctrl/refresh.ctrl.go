package authctrl

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"pentag.kr/distimer/db"
	"pentag.kr/distimer/ent"
	"pentag.kr/distimer/ent/session"
	"pentag.kr/distimer/utils/crypt"
	"pentag.kr/distimer/utils/dto"
	"pentag.kr/distimer/utils/logger"
)

func Refresh(c *fiber.Ctx) error {
	data := new(refreshTokenDTO)
	if err := dto.Bind(c, data); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err,
		})
	}
	dbConn := db.GetDBClient()
	refreshToken := uuid.MustParse(data.RefreshToken)

	sessionObj, err := dbConn.Session.Query().Where(session.RefreshToken(refreshToken)).Only(context.Background())
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

	owner, err := sessionObj.QueryUser().Only(context.Background())
	if err != nil {
		logger.CtxError(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	newAccessToken := crypt.NewJWT(owner.ID, owner.TermsAgreed)
	newRefreshToken := uuid.New()

	err = dbConn.Session.UpdateOne(sessionObj).SetRefreshToken(newRefreshToken).Exec(context.Background())
	if err != nil {
		logger.CtxError(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	return c.JSON(accessInfoDTO{
		SessionID:    sessionObj.ID.String(),
		UserID:       owner.ID.String(),
		Name:         owner.Name,
		RefreshToken: newRefreshToken.String(),
		AccessToken:  newAccessToken,
	})
}
