package userctrl

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"pentag.kr/distimer/db"
	"pentag.kr/distimer/middlewares"
	"pentag.kr/distimer/utils/logger"
)

type myUserInfoRes struct {
	UserID      string `json:"user_id"`
	Name        string `json:"name"`
	TermsAgreed bool   `json:"terms_agreed"`
	CreatedAt   string `json:"created_at"`
}

// @Summary Get My User Info
// @Tags User
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} myUserInfoRes
// @Router /user [get]
func GetMyUserInfo(c *fiber.Ctx) error {
	userID := middlewares.GetUserIDFromMiddleware(c)

	dbConn := db.GetDBClient()

	user, err := dbConn.User.Get(context.Background(), userID)
	if err != nil {
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}
	return c.JSON(myUserInfoRes{
		UserID:      user.ID.String(),
		Name:        user.Name,
		TermsAgreed: user.TermsAgreed,
		CreatedAt:   user.CreatedAt.Format(time.RFC3339),
	})
}
