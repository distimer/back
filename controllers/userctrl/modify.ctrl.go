package userctrl

import (
	"context"
	"unicode/utf8"

	"github.com/gofiber/fiber/v2"
	"pentag.kr/distimer/db"
	"pentag.kr/distimer/middlewares"
	"pentag.kr/distimer/utils/dto"
	"pentag.kr/distimer/utils/logger"
)

type modifyUserInfoReq struct {
	Name        string `json:"name" validate:"required" example:"name between 1 and 20"`
	TermsAgreed bool   `json:"terms_agreed"`
}
type modifyUserInfoRes struct {
	UserID      string `json:"user_id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	TermsAgreed bool   `json:"terms_agreed" validate:"required"`
}

// @Summary Modify User Info
// @Tags User
// @Accept json
// @Produce json
// @Security Bearer
// @Param request body modifyUserInfoReq true "modifyUserInfoReq"
// @Success 200 {object} modifyUserInfoRes
// @Router /user [put]
func ModifyUserInfo(c *fiber.Ctx) error {
	data := new(modifyUserInfoReq)
	if err := dto.Bind(c, data); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err,
		})
	}
	if utf8.RuneCountInString(data.Name) < 1 || utf8.RuneCountInString(data.Name) > 20 {
		return c.Status(400).JSON(fiber.Map{
			"error": "Name length should be between 1 and 20",
		})
	}

	userID := middlewares.GetUserIDFromMiddleware(c)

	dbConn := db.GetDBClient()

	_, err := dbConn.User.UpdateOneID(userID).SetName(data.Name).SetTermsAgreed(data.TermsAgreed).Save(context.Background())
	if err != nil {
		logger.CtxError(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	return c.JSON(modifyUserInfoRes{
		UserID: userID.String(),
		Name:   data.Name,
	})
}
