package studylogctrl

import (
	"context"
	"time"

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

	logObj, err := dbConn.StudyLog.Query().Where(studylog.And(studylog.ID(studylogID), studylog.HasUserWith(user.ID(userID)))).WithSubject().WithSharedGroup().Only(context.Background())

	if err != nil {
		if ent.IsNotFound(err) {
			return c.Status(404).JSON(fiber.Map{
				"error": "StudyLog not found",
			})
		}
		logger.CtxError(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}
	result := myStudyLogDTO{
		ID:        logObj.ID.String(),
		SubjectID: logObj.Edges.Subject.ID.String(),
		StartAt:   logObj.StartAt.Format(time.RFC3339),
		EndAt:     logObj.EndAt.Format(time.RFC3339),
		Content:   logObj.Content,
		GroupsToShare: func() []string {
			groups := make([]string, len(logObj.Edges.SharedGroup))
			for i, v := range logObj.Edges.SharedGroup {
				groups[i] = v.ID.String()
			}
			return groups
		}(),
	}
	return c.JSON(result)
}
