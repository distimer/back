package studylogctrl

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"pentag.kr/distimer/db"
	"pentag.kr/distimer/ent"
	"pentag.kr/distimer/ent/studylog"
	"pentag.kr/distimer/ent/user"
	"pentag.kr/distimer/middlewares"
	"pentag.kr/distimer/utils/logger"
)

// @Summary Get StudyLogs with Term
// @Tags StudyLog
// @Accept json
// @Produce json
// @Security Bearer
// @Param start_date query string true "2006-01-02"
// @Param end_date query string true "2006-01-03"
// @Success 200 {array} myStudyLogDTO
// @Failure 400
// @Failure 500
// @Router /studylog/term [get]
func GetByTerm(c *fiber.Ctx) error {
	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")

	startDate, err := time.Parse(time.RFC3339, startDateStr+"T00:00:00.000+09:00")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid start_date",
		})
	}
	if startDate.After(time.Now()) {
		return c.Status(400).JSON(fiber.Map{
			"error": "date should be before today",
		})
	}

	endDate, err := time.Parse(time.RFC3339, endDateStr+"T00:00:00.000+09:00")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid end_date",
		})

	}
	if endDate.After(time.Now()) {
		return c.Status(400).JSON(fiber.Map{
			"error": "date should be before today",
		})
	}

	if startDate.After(endDate) {
		return c.Status(400).JSON(fiber.Map{
			"error": "start_date should be before end_date",
		})
	}
	// term date
	termDate := int(endDate.Sub(startDate).Hours()/24) + 1
	if termDate > maxTermDateRange {
		return c.Status(400).JSON(fiber.Map{
			"error": "The term date range should be less than 31 days",
		})
	}

	userID := middlewares.GetUserIDFromMiddleware(c)
	dbConn := db.GetDBClient()

	studylogs, err := dbConn.StudyLog.Query().
		Where(
			studylog.And(
				studylog.HasUserWith(user.ID(userID)),
				studylog.And(
					studylog.StartAtLT(endDate.AddDate(0, 0, 1)),
					studylog.EndAtGT(startDate),
				)),
		).Order(ent.Asc("start_at")).WithSubject().All(context.Background())
	if err != nil {
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	result := make([]myStudyLogDTO, len(studylogs))
	for i, log := range studylogs {
		result[i] = myStudyLogDTO{
			ID:        log.ID.String(),
			SubjectID: log.Edges.Subject.ID.String(),
			StartAt:   log.StartAt.Format(time.RFC3339),
			EndAt:     log.EndAt.Format(time.RFC3339),
			Content:   log.Content,
			GroupsToShare: func() []string {
				groups := make([]string, len(log.Edges.SharedGroup))
				for i, v := range log.Edges.SharedGroup {
					groups[i] = v.ID.String()
				}
				return groups
			}(),
		}
	}
	return c.JSON(result)
}
