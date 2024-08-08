package studylogctrl

import (
	"context"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"pentag.kr/distimer/db"
	"pentag.kr/distimer/ent"
	"pentag.kr/distimer/ent/studylog"
	"pentag.kr/distimer/ent/user"
	"pentag.kr/distimer/middlewares"
	"pentag.kr/distimer/utils/logger"
)

const (
	maxCount = 20
)

// @Summary Get All My Study Logs
// @Tags StudyLog
// @Accept json
// @Produce json
// @Security Bearer
// @Param count query int true "count"
// @Param offset query int true "offset"
// @Success 200 {array} myStudyLogDTO
// @Router /studylog [get]
func GetAllMyStudyLogs(c *fiber.Ctx) error {
	userID := middlewares.GetUserIDFromMiddleware(c)
	countStr := c.Query("count")
	offsetStr := c.Query("offset")

	count, err := strconv.Atoi(countStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid count",
		})
	} else if count > maxCount || count < 1 {
		return c.Status(400).JSON(fiber.Map{
			"error": "count should be between 1 and 20",
		})
	}
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid offset",
		})
	} else if offset < 0 {
		return c.Status(400).JSON(fiber.Map{
			"error": "offset should be greater than or equal to 0",
		})
	}

	// Fetch study logs from database
	dbConn := db.GetDBClient()
	studyLogs, err := dbConn.StudyLog.Query().
		Where(studylog.HasUserWith(user.ID(userID))).
		WithSharedGroup().
		WithSubject().
		Order(ent.Desc("start_at")).
		Limit(count).
		Offset(offset).
		All(context.Background())
	if err != nil {
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}
	result := make([]myStudyLogDTO, len(studyLogs))
	for i, studyLog := range studyLogs {
		result[i] = myStudyLogDTO{
			ID:        studyLog.ID.String(),
			SubjectID: studyLog.Edges.Subject.ID.String(),
			StartAt:   studyLog.StartAt.Format("2006-01-02 15:04:05"),
			EndAt:     studyLog.EndAt.Format("2006-01-02 15:04:05"),
			Content:   studyLog.Content,
			GroupsToShare: func() []string {
				groups := make([]string, len(studyLog.Edges.SharedGroup))
				for i, group := range studyLog.Edges.SharedGroup {
					groups[i] = group.ID.String()
				}
				return groups
			}(),
		}
	}
	return c.JSON(result)
}
