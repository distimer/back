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

type logoutTokenReq struct {
	dto.BaseDTO
	RefreshToken string `json:"refresh_token" validate:"required,uuid"`
}

// @Summary Logout
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body authctrl.logoutTokenReq true "logoutTokenReq"
// @Success 204
// @Router /auth/logout [delete]
func Logout(c *fiber.Ctx) error {
	data := new(logoutTokenReq)
	if err := dto.Bind(c, data); err != nil {
		return err
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
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	return c.SendStatus(204)
}
