package subjectctrl

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"pentag.kr/distimer/db"
	"pentag.kr/distimer/ent"
	"pentag.kr/distimer/ent/studylog"
	"pentag.kr/distimer/ent/subject"
	"pentag.kr/distimer/ent/user"
	"pentag.kr/distimer/middlewares"
)

type getStudyLogRes struct {
	StudyLogs []*ent.StudyLog `json:"study_logs"`
}

// @Summary Get Study Log with Subject
// @Tags Subject
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "Subject ID"
// @Success 200 {object} getStudyLogRes
// @Router /subject/studylog/{id} [get]
func GetStudyLogWithSubject(c *fiber.Ctx) error {
	subjectIDStr := c.Params("id")

	subjectID, err := uuid.Parse(subjectIDStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid subject ID",
		})
	}
	userID := middlewares.GetUserIDFromMiddleware(c)

	dbConn := db.GetDBClient()

	studyLogs, err := dbConn.StudyLog.Query().Where(studylog.And(studylog.HasUserWith(user.ID(userID)), studylog.HasSubjectWith(subject.ID(subjectID)))).Order(ent.Desc("created_at")).All(context.Background())

	if err != nil {
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
