package groupctrl

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"pentag.kr/distimer/db"
	"pentag.kr/distimer/ent/affiliation"
	"pentag.kr/distimer/ent/group"
	"pentag.kr/distimer/ent/studylog"
	"pentag.kr/distimer/ent/timer"
	"pentag.kr/distimer/ent/user"
	"pentag.kr/distimer/middlewares"
	"pentag.kr/distimer/utils/logger"
)

// @Summary Expel Member
// @Tags Group
// @Accept json
// @Produce json
// @Security Bearer
// @Param groupID path string true "group id"
// @Param memberID path string true "member id"
// @Success 204
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /group/member/{groupID}/{memberID} [delete]
func ExpelMember(c *fiber.Ctx) error {
	groupIDStr := c.Params("group_id")
	memberIDStr := c.Params("member_id")

	groupID, err := uuid.Parse(groupIDStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid group id",
		})
	}
	memberID, err := uuid.Parse(memberIDStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid member id",
		})
	}

	userID := middlewares.GetUserIDFromMiddleware(c)
	if userID == memberID {
		return c.Status(400).JSON(fiber.Map{
			"error": "You cannot modify yourself",
		})
	}

	dbConn := db.GetDBClient()

	exist, err := dbConn.Group.Query().Where(group.And(group.ID(groupID), group.HasOwnerWith(user.ID(userID)))).Exist(context.Background())
	if err != nil {
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	} else if !exist {
		return c.Status(404).JSON(fiber.Map{
			"error": "Group not found or you are not the owner",
		})
	}

	// delete affiliation
	deletedCount, err := dbConn.Affiliation.Delete().Where(affiliation.And(affiliation.GroupID(groupID), affiliation.UserID(memberID))).Exec(context.Background())
	if err != nil {
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}
	if deletedCount == 0 {
		return c.Status(404).JSON(fiber.Map{
			"error": "User not found in the group",
		})
	}

	_, err = dbConn.StudyLog.Update().Where(
		studylog.And(
			studylog.HasSharedGroupWith(group.ID(groupID)),
			studylog.HasUserWith(user.ID(memberID)),
		)).
		RemoveSharedGroupIDs(groupID).
		Save(context.Background())
	if err != nil {
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	_, err = dbConn.Timer.Update().Where(
		timer.And(
			timer.HasSharedGroupWith(group.ID(groupID)),
			timer.HasUserWith(user.ID(memberID)),
		)).
		RemoveSharedGroupIDs(groupID).
		Save(context.Background())
	if err != nil {
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}
	return c.SendStatus(204)
}
