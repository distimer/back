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

type refreshTokenReq struct {
	RefreshToken string `json:"refresh_token" validate:"required,uuid"`
}

type refreshTokenRes struct {
	RefreshToken string `json:"refresh_token" validate:"required,uuid"`
	AccessToken  string `json:"access_token" validate:"required"`
}

// @Summary Refresh Token
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body authctrl.refreshTokenReq true "refreshTokenReq"
// @Success 200 {object} refreshTokenRes
// @Router /auth/refresh [post]
func Refresh(c *fiber.Ctx) error {
	data := new(refreshTokenReq)
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
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	owner, err := refreshTokenObj.QueryUser().Only(context.Background())
	if err != nil {
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	newAccessToken := crypt.NewJWT(owner.ID, owner.TermsAgreed)
	newRefrshToken := uuid.New()

	err = dbConn.RefreshToken.DeleteOne(refreshTokenObj).Exec(context.Background())
	if err != nil {
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	_, err = dbConn.RefreshToken.Create().
		SetID(newRefrshToken).
		SetUser(owner).
		Save(context.Background())
	if err != nil {
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}
	return c.JSON(refreshTokenRes{
		RefreshToken: newRefrshToken.String(),
		AccessToken:  newAccessToken,
	})
}
