package invitectrl

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"pentag.kr/distimer/db"
	"pentag.kr/distimer/ent"
	"pentag.kr/distimer/ent/affiliation"
	"pentag.kr/distimer/ent/group"
	"pentag.kr/distimer/ent/invitecode"
	"pentag.kr/distimer/middlewares"
	"pentag.kr/distimer/utils/logger"
)

// @Summary Get Invite Code List
// @Tags Invite
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "group id"
// @Success 200 {array} string
// @Router /invite/group/{id} [get]
func GetInviteCodeList(c *fiber.Ctx) error {
	groupIDStr := c.Params("id")
	groupID, err := uuid.Parse(groupIDStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid group id",
		})
	}

	userID := middlewares.GetUserIDFromMiddleware(c)

	dbConn := db.GetDBClient()

	affiliationObj, err := dbConn.Affiliation.Query().Where(affiliation.And(affiliation.GroupID(groupID), affiliation.UserID(userID))).Only(context.Background())
	if err != nil {
		if ent.IsNotFound(err) {
			return c.Status(404).JSON(fiber.Map{
				"error": "Group is not exist , or you are not the member of the group",
			})
		}
		logger.CtxError(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}
	if affiliationObj.Role != 2 {
		groupObj, err := dbConn.Group.Query().Where(group.ID(groupID)).Only(context.Background())
		if err != nil {
			logger.CtxError(c, err)
			return c.Status(500).JSON(fiber.Map{
				"error": "Internal server error",
			})
		} else if affiliationObj.Role < groupObj.InvitePolicy {
			return c.Status(403).JSON(fiber.Map{
				"error": "You are not allowed to invite",
			})
		}
	}

	inviteCodeList, err := dbConn.InviteCode.Query().
		Where(invitecode.HasGroupWith(group.ID(groupID))).
		Select(invitecode.FieldCode).
		Strings(context.Background())
	if err != nil {
		logger.CtxError(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}
	if len(inviteCodeList) == 0 {
		inviteCodeList = []string{}
	}
	return c.JSON(inviteCodeList)
}
