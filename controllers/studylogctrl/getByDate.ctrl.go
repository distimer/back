package studylogctrl

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"pentag.kr/distimer/db"
	"pentag.kr/distimer/ent/studylog"
	"pentag.kr/distimer/ent/user"
	"pentag.kr/distimer/middlewares"
)

// @Summary Get StudyLog by Date
// @Tags StudyLog
// @Accept json
// @Produce json
// @Security Bearer
// @Param date query string false "2006-01-02"
// @Success 200 {array} myStudyLogDTO
// @Failure 400
// @Failure 500
// @Router /studylog/date [get]
func GetByDate(c *fiber.Ctx) error {
	dateStr := c.Query("date", "")
	date, err := time.Parse(time.RFC3339, dateStr+"T00:00:00.000+09:00")
	if err != nil {
		date = time.Now()
		date = date.Truncate(24 * time.Hour)
		date = date.Add(-9 * time.Hour)
	}
	if date.After(time.Now()) {
		return c.Status(400).JSON(fiber.Map{
			"error": "date should be before today",
		})
	}

	userID := middlewares.GetUserIDFromMiddleware(c)

	dbConn := db.GetDBClient()

	studylogList, err := dbConn.StudyLog.Query().
		Where(
			studylog.And(studylog.And(
				studylog.StartAtLTE(date.AddDate(0, 0, 1)),
				studylog.EndAtGTE(date),
			), studylog.HasUserWith(user.ID(userID))),
		).
		WithSubject().
		WithSharedGroup().
		All(context.Background())

	result := make([]myStudyLogDTO, len(studylogList))
	for i, v := range studylogList {
		result[i] = myStudyLogDTO{
			ID:        v.ID.String(),
			SubjectID: v.Edges.Subject.ID.String(),
			StartAt:   v.StartAt.Format(time.RFC3339),
			EndAt:     v.EndAt.Format(time.RFC3339),
			Content:   v.Content,
			GroupsToShare: func() []string {
				groups := make([]string, len(v.Edges.SharedGroup))
				for i, v := range v.Edges.SharedGroup {
					groups[i] = v.ID.String()
				}
				return groups
			}(),
		}
	}
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}
	return c.JSON(result)

}
