package studylogctrl

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"pentag.kr/distimer/db"
	"pentag.kr/distimer/ent/studylog"
	"pentag.kr/distimer/ent/user"
	"pentag.kr/distimer/middlewares"
	"pentag.kr/distimer/utils/logger"
)

// @Summary Delete StudyLog
// @Tags StudyLog
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "studylog id"
// @Success 204
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /studylog/{id} [delete]
func DeleteStudyLog(c *fiber.Ctx) error {
	studylogIDStr := c.Params("id")
	studylogID, err := uuid.Parse(studylogIDStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid studylog id",
		})
	}

	userID := middlewares.GetUserIDFromMiddleware(c)

	dbConn := db.GetDBClient()

	_, err = dbConn.StudyLog.Delete().Where(
		studylog.And(
			studylog.ID(studylogID),
			studylog.HasUserWith(user.ID(userID)),
		),
	).Exec(context.Background())
	if err != nil {
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	return c.SendStatus(204)
}
