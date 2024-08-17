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

const (
	maxTermDateRange = 62
)

type dailySubjectLog struct {
	SubjectID  string `json:"subject_id" validate:"required"`
	CategoryID string `json:"category_id" validate:"required"`
	StudyTime  int    `json:"study_time" validate:"required"`
}

type dailyStudyLog struct {
	Date string            `json:"date" validate:"required"`
	Log  []dailySubjectLog `json:"log" validate:"required"`
}

// @Summary Get Statistics with Term
// @Tags StudyLog
// @Accept json
// @Produce json
// @Security Bearer
// @Param start_date query string true "2006-01-02"
// @Param end_date query string true "2006-01-03"
// @Success 200 {array} dailyStudyLog
// @Failure 400
// @Failure 500
// @Router /studylog/statistics/term [get]
func GetStatisticsByTerm(c *fiber.Ctx) error {
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
			"error": "The term date range should be less than 62 days",
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
		logger.CtxError(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	// inject category edge
	for _, log := range studylogs {
		log.Edges.Subject.Edges.Category, err = log.Edges.Subject.QueryCategory().Only(context.Background())
		if err != nil {
			logger.CtxError(c, err)
			return c.Status(500).JSON(fiber.Map{
				"error": "Internal server error",
			})
		}
	}

	// create array of daily study logs
	dailyStudyLogs := make([]dailyStudyLog, termDate)
	dateCounter := startDate

	for i := range dailyStudyLogs {
		dailyStudyLogs[i].Date = dateCounter.Format("2006-01-02")
		dailyStudyLogs[i].Log = make([]dailySubjectLog, 0)
		dateCounter = dateCounter.AddDate(0, 0, 1)
	}

	for _, log := range studylogs {
		startDateIndex := int(log.StartAt.Sub(startDate).Hours() / 24)
		endDateIndex := int(log.EndAt.Sub(startDate).Hours() / 24)
		if startDateIndex != endDateIndex {
			// log is separated by date
			dailyStudyLogs[startDateIndex].Log = append(
				dailyStudyLogs[startDateIndex].Log,
				dailySubjectLog{
					SubjectID:  log.Edges.Subject.ID.String(),
					CategoryID: log.Edges.Subject.Edges.Category.ID.String(),
					StudyTime:  int(startDate.AddDate(0, 0, startDateIndex+1).Sub(log.StartAt).Seconds()),
				},
			)
			if startDateIndex < termDate-1 {
				dailyStudyLogs[endDateIndex].Log = append(
					dailyStudyLogs[endDateIndex].Log,
					dailySubjectLog{
						SubjectID:  log.Edges.Subject.ID.String(),
						CategoryID: log.Edges.Subject.Edges.Category.ID.String(),
						StudyTime:  int(log.EndAt.Sub(startDate.AddDate(0, 0, endDateIndex)).Seconds()),
					},
				)
			}
		} else {
			dailyStudyLogs[startDateIndex].Log = append(
				dailyStudyLogs[startDateIndex].Log,
				dailySubjectLog{
					SubjectID:  log.Edges.Subject.ID.String(),
					CategoryID: log.Edges.Subject.Edges.Category.ID.String(),
					StudyTime:  int(log.EndAt.Sub(log.StartAt).Seconds()),
				},
			)
		}
	}
	return c.JSON(dailyStudyLogs)
}
