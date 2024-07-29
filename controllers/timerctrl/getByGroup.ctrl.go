package timerctrl

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"pentag.kr/distimer/db"
	"pentag.kr/distimer/ent"
	"pentag.kr/distimer/ent/affiliation"
	"pentag.kr/distimer/middlewares"
	"pentag.kr/distimer/utils/logger"
)

// @Summary Get Timer by Group
// @Tags Timer
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "Group ID"
// @Success 200 {array} timerDTO
// @Failure 400
// @Failure 403
// @Failure 404
// @Failure 500
// @Router /timer/group/{id} [get]
func GetTimerByGroup(c *fiber.Ctx) error {

	groupIDStr := c.Params("id")
	groupID, err := uuid.Parse(groupIDStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid group ID",
		})
	}

	userID := middlewares.GetUserIDFromMiddleware(c)

	dbConn := db.GetDBClient()

	affiliationObj, err := dbConn.Affiliation.Query().
		Where(affiliation.GroupID(groupID), affiliation.UserID(userID)).
		WithGroup().
		Only(context.Background())
	if err != nil {
		if ent.IsNotFound(err) {
			return c.Status(404).JSON(fiber.Map{
				"error": "Group not found or not affiliated",
			})
		}
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}
	if affiliationObj.Edges.Group.RevealPolicy > affiliationObj.Role {
		return c.Status(403).JSON(fiber.Map{
			"error": "Not Enough Permission",
		})
	}

	timers, err := affiliationObj.Edges.Group.QuerySharedTimer().All(context.Background())
	if err != nil {
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}
	result := make([]timerDTO, len(timers))
	for i, timer := range timers {
		result[i] = timerDTO{
			ID:        timer.ID.String(),
			SubjectID: timer.SubjectID.String(),
			Content:   timer.Content,
			StartAt:   timer.StartAt.Format(time.RFC3339),
		}
	}
	return c.JSON(result)
}