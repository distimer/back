package groupctrl

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"pentag.kr/distimer/db"
	"pentag.kr/distimer/ent"
	"pentag.kr/distimer/ent/affiliation"
	"pentag.kr/distimer/ent/group"
	"pentag.kr/distimer/ent/studylog"
	"pentag.kr/distimer/ent/timer"
	"pentag.kr/distimer/ent/user"
	"pentag.kr/distimer/middlewares"
	"pentag.kr/distimer/utils/logger"
)

// @Summary Quit Group
// @Tags Group
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "group id"
// @Success 204
// @Failure 400
// @Failure 403
// @Failure 404
// @Failure 500
// @Router /group/quit/{id} [delete]
func QuitGroup(c *fiber.Ctx) error {
	groupIDStr := c.Params("id")
	groupID, err := uuid.Parse(groupIDStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid group ID",
		})
	}
	userID := middlewares.GetUserIDFromMiddleware(c)

	dbConn := db.GetDBClient()

	affiliationObj, err := dbConn.Affiliation.Query().Where(
		affiliation.And(
			affiliation.GroupID(groupID),
			affiliation.UserID(userID),
		),
	).Only(c.Context())
	if err != nil {
		if ent.IsNotFound(err) {
			return c.Status(404).JSON(fiber.Map{
				"error": "Affiliation not found",
			})
		}
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	} else if affiliationObj.Role == 2 {
		return c.Status(403).JSON(fiber.Map{
			"error": "Owner cannot quit group",
		})
	}

	// get user timer and delete group from shared group
	_, err = dbConn.Timer.Update().Where(
		timer.HasUserWith(user.ID(userID)),
		timer.HasSharedGroupWith(group.ID(groupID)),
	).RemoveSharedGroupIDs(groupID).Save(c.Context())
	if err != nil {
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	// delete group from shared group
	_, err = dbConn.StudyLog.Update().Where(
		studylog.And(
			studylog.HasUserWith(user.ID(userID)),
			studylog.HasSharedGroupWith(group.ID(groupID)),
		),
	).RemoveSharedGroupIDs(groupID).Save(context.Background())
	if err != nil {
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	// delete affiliation
	_, err = dbConn.Affiliation.Delete().Where(affiliation.And(
		affiliation.GroupID(groupID),
		affiliation.UserID(userID),
	)).Exec(context.Background())
	if err != nil {
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	return c.SendStatus(204)
}
