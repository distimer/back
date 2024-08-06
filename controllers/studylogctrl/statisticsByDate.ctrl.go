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

// @Summary Get statistics of study logs with date
// @Tags StudyLog
// @Accept json
// @Produce json
// @Security Bearer
// @Param date query string true "2006-01-02"
// @Success 200 {array} dailySubjectLog
// @Failure 400
// @Failure 500
// @Router /studylog/statistics/date [get]
func GetStatisticsWithDate(c *fiber.Ctx) error {

	dateStr := c.Query("date")
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid date format",
		})
	}

	if date.After(time.Now()) {
		return c.Status(400).JSON(fiber.Map{
			"error": "date should be before today",
		})
	}

	userID := middlewares.GetUserIDFromMiddleware(c)
	dbConn := db.GetDBClient()

	// get studylog

	logs, err := dbConn.StudyLog.Query().
		Where(
			studylog.And(
				studylog.HasUserWith(user.ID(userID)),
				studylog.And(
					studylog.StartAtLT(date.AddDate(0, 0, 1)),
					studylog.EndAtGT(date),
				)),
		).Order(ent.Asc("start_at")).WithSubject().All(context.Background())
	if err != nil {
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	for _, log := range logs {
		log.Edges.Subject.Edges.Category, err = log.Edges.Subject.QueryCategory().Only(context.Background())
		if err != nil {
			logger.Error(c, err)
			return c.Status(500).JSON(fiber.Map{
				"error": "Internal server error",
			})
		}
	}

	result := []dailySubjectLog{}
	for _, log := range logs {
		if log.StartAt.Before(date) {
			result = append(result, dailySubjectLog{
				SubjectID:  log.Edges.Subject.ID.String(),
				StudyTime:  int(log.EndAt.Sub(date).Seconds()),
				CategoryID: log.Edges.Subject.Edges.Category.ID.String(),
			})
		} else if log.EndAt.After(date.AddDate(0, 0, 1)) {
			result = append(result, dailySubjectLog{
				SubjectID:  log.Edges.Subject.ID.String(),
				StudyTime:  int(date.AddDate(0, 0, 1).Sub(log.StartAt).Seconds()),
				CategoryID: log.Edges.Subject.Edges.Category.ID.String(),
			})
		} else {
			result = append(result, dailySubjectLog{
				SubjectID:  log.Edges.Subject.ID.String(),
				StudyTime:  int(log.EndAt.Sub(log.StartAt).Seconds()),
				CategoryID: log.Edges.Subject.Edges.Category.ID.String(),
			})
		}
	}
	return c.JSON(result)
}
