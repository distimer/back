package groupctrl

import (
	"context"
	"time"
	"unicode/utf8"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"pentag.kr/distimer/db"
	"pentag.kr/distimer/ent"
	"pentag.kr/distimer/ent/affiliation"
	"pentag.kr/distimer/middlewares"
	"pentag.kr/distimer/utils/dto"
	"pentag.kr/distimer/utils/logger"
)

type modifyNicknameReq struct {
	Nickname string `json:"nickname" validate:"required" example:"nickname between 1 and 20"`
}

// @Summary Modify Nickname
// @Tags Group
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "group id"
// @Param request body modifyNicknameReq true "modifyNicknameReq"
// @Success 200 {object} AffiliationDTO
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /group/nickname/{id} [patch]
func ModifyNickname(c *fiber.Ctx) error {
	groupIDStr := c.Params("id")
	groupID, err := uuid.Parse(groupIDStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid group id",
		})
	}
	data := new(modifyNicknameReq)
	if err := dto.Bind(c, data); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err,
		})
	}
	if utf8.RuneCountInString(data.Nickname) < 1 || utf8.RuneCountInString(data.Nickname) > 20 {
		return c.Status(400).JSON(fiber.Map{
			"error": "Nickname length should be between 1 and 20",
		})
	}

	userID := middlewares.GetUserIDFromMiddleware(c)
	dbConn := db.GetDBClient()

	affiliationObj, err := dbConn.Affiliation.Query().Where(affiliation.And(affiliation.GroupID(groupID), affiliation.UserID(userID))).Only(context.Background())
	if err != nil {
		if ent.IsNotFound(err) {
			return c.Status(404).JSON(fiber.Map{
				"error": "You are not a member of this group",
			})
		}
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	affiliationObj, err = affiliationObj.Update().SetNickname(data.Nickname).Save(context.Background())
	if err != nil {
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	return c.JSON(AffiliationDTO{
		GroupID:  groupIDStr,
		UserID:   userID.String(),
		Nickname: affiliationObj.Nickname,
		Role:     affiliationObj.Role,
		JoinedAt: affiliationObj.JoinedAt.Format(time.RFC3339),
	})
}
