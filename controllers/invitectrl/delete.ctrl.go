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

// @Summary Delete Invite Code
// @Tags Invite
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "group id"
// @Param code path string true "invite code"
// @Success 204
// @Router /invite/group/{id}/{code} [delete]
func DeleteInviteCode(c *fiber.Ctx) error {
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
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}
	if affiliationObj.Role != 2 {
		groupObj, err := dbConn.Group.Query().Where(group.ID(groupID)).Only(context.Background())
		if err != nil {
			logger.Error(c, err)
			return c.Status(500).JSON(fiber.Map{
				"error": "Internal server error",
			})
		} else if affiliationObj.Role < groupObj.InvitePolicy {
			return c.Status(403).JSON(fiber.Map{
				"error": "You are not allowed to invite",
			})
		}
	}

	code := c.Params("code")

	_, err = dbConn.InviteCode.Delete().
		Where(invitecode.And(invitecode.HasGroupWith(group.ID(groupID)), invitecode.Code(code))).
		Exec(context.Background())
	if err != nil {
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}
	return c.SendStatus(204)
}
