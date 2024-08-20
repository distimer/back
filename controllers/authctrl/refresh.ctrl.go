package authctrl

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"pentag.kr/distimer/db"
	"pentag.kr/distimer/ent"
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

	refreshTokenObj, err := dbConn.RefreshToken.Get(context.Background(), refreshToken)
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

	owner, err := refreshTokenObj.QueryUser().Only(context.Background())
	if err != nil {
		logger.CtxError(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	newAccessToken := crypt.NewJWT(owner.ID, owner.TermsAgreed)
	newRefrshToken := uuid.New()

	err = dbConn.RefreshToken.DeleteOne(refreshTokenObj).Exec(context.Background())
	if err != nil {
		logger.CtxError(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	_, err = dbConn.RefreshToken.Create().
		SetID(newRefrshToken).
		SetUser(owner).
		Save(context.Background())
	if err != nil {
		logger.CtxError(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}
	return c.JSON(accessInfoDTO{
		UserID:       owner.ID.String(),
		Name:         owner.Name,
		RefreshToken: newRefrshToken.String(),
		AccessToken:  newAccessToken,
	})
}
