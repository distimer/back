package studylogctrl

import (
	"context"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"pentag.kr/distimer/db"
	"pentag.kr/distimer/ent"
	"pentag.kr/distimer/ent/studylog"
	"pentag.kr/distimer/ent/subject"
	"pentag.kr/distimer/ent/user"
	"pentag.kr/distimer/middlewares"
	"pentag.kr/distimer/utils/logger"
)

type getStudyLogRes struct {
	StudyLogs []*ent.StudyLog `json:"study_logs"`
}

// @Summary Get Study Log with Subject
// @Tags StudyLog
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "Subject ID"
// @Param count query int true "count"
// @Param offset query int true "offset"
// @Success 200 {object} getStudyLogRes
// @Router /studylog/subject/{id} [get]
func GetStudyLogWithSubject(c *fiber.Ctx) error {
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

	subjectIDStr := c.Params("id")

	subjectID, err := uuid.Parse(subjectIDStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid subject ID",
		})
	}
	userID := middlewares.GetUserIDFromMiddleware(c)

	dbConn := db.GetDBClient()

	studyLogs, err := dbConn.StudyLog.Query().
		Where(studylog.And(studylog.HasUserWith(user.ID(userID)), studylog.HasSubjectWith(subject.ID(subjectID)))).
		Order(ent.Desc("created_at")).
		Limit(count).
		Offset(offset).
		All(context.Background())

	if err != nil {
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}
	if len(studyLogs) == 0 {
		studyLogs = []*ent.StudyLog{}
	}
	return c.JSON(getStudyLogRes{
		StudyLogs: studyLogs,
	})
}
