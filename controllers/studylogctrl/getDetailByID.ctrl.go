package studylogctrl

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"pentag.kr/distimer/db"
	"pentag.kr/distimer/ent"
	"pentag.kr/distimer/ent/studylog"
	"pentag.kr/distimer/ent/user"
	"pentag.kr/distimer/middlewares"
	"pentag.kr/distimer/utils/logger"
)

// @Summary Get StudyLog Detail by ID
// @Tags StudyLog
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "studylog id"
// @Success 200 {object} myStudyLogDTO
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /studylog/detail/{id} [get]
func GetDetailByID(c *fiber.Ctx) error {
	studylogIDStr := c.Params("id")
	studylogID, err := uuid.Parse(studylogIDStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid studylog id",
		})
	}
	userID := middlewares.GetUserIDFromMiddleware(c)
	dbConn := db.GetDBClient()

	logObj, err := dbConn.StudyLog.Query().Where(studylog.And(studylog.ID(studylogID), studylog.HasUserWith(user.ID(userID)))).WithSubject().Only(context.Background())

	if err != nil {
		if ent.IsNotFound(err) {
			return c.Status(404).JSON(fiber.Map{
				"error": "StudyLog not found",
			})
		}
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}
	result := myStudyLogDTO{
		ID:        logObj.ID.String(),
		SubjectID: logObj.Edges.Subject.ID.String(),
		StartAt:   logObj.StartAt.Format("2006-01-02T15:04:05"),
		EndAt:     logObj.EndAt.Format("2006-01-02T15:04:05"),
		Content:   logObj.Content,
	}
	return c.JSON(result)
}
