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
	maxTermDateRange = 31
)

type dailySubjectLog struct {
	SubjectID  string `json:"subject_id"`
	CategoryID string `json:"category_id"`
	StudyTime  int    `json:"study_time"`
}

type dailyStudyLog struct {
	Date string            `json:"date"`
	Log  []dailySubjectLog `json:"log"`
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
func GetStatisticsWithTerm(c *fiber.Ctx) error {
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

	// inject category edge
	for _, log := range studylogs {
		log.Edges.Subject.Edges.Category, err = log.Edges.Subject.QueryCategory().Only(context.Background())
		if err != nil {
			logger.Error(c, err)
			return c.Status(500).JSON(fiber.Map{
				"error": "Internal server error",
			})
		}
	}

	// Calculate the total study time

	if len(studylogs) == 0 {
		return c.JSON([]dailyStudyLog{})
	}

	// create array of daily study logs
	dailyStudyLogs := make([]dailyStudyLog, termDate)
	dateCounter := 0
	logCounter := 0
	date := startDate
	for dateCounter < termDate {

		firstLog := studylogs[logCounter]

		// check if the log is separated by date
		if firstLog.StartAt.Day() < date.Day() {
			dailyStudyLogs[dateCounter].Log = append(
				dailyStudyLogs[dateCounter].Log,
				dailySubjectLog{
					SubjectID:  firstLog.Edges.Subject.ID.String(),
					CategoryID: firstLog.Edges.Subject.Edges.Category.ID.String(),
					StudyTime:  int(firstLog.EndAt.Sub(date).Seconds()),
				},
			)
			logCounter++
		}

		flag := true
		for flag && logCounter != len(studylogs) && studylogs[logCounter].StartAt.Day() <= date.Day() {
			log := studylogs[logCounter]
			if log.EndAt.Day() > date.Day() {
				flag = false
				dailyStudyLogs[dateCounter].Log = append(
					dailyStudyLogs[dateCounter].Log,
					dailySubjectLog{
						SubjectID:  log.Edges.Subject.ID.String(),
						CategoryID: log.Edges.Subject.Edges.Category.ID.String(),
						StudyTime:  int(date.AddDate(0, 0, 1).Sub(log.StartAt).Seconds()),
					},
				)
				break
			} else {
				dailyStudyLogs[dateCounter].Log = append(
					dailyStudyLogs[dateCounter].Log,
					dailySubjectLog{
						SubjectID:  log.Edges.Subject.ID.String(),
						CategoryID: log.Edges.Subject.Edges.Category.ID.String(),
						StudyTime:  int(log.EndAt.Sub(log.StartAt).Seconds()),
					},
				)
				logCounter++
			}
		}
		dailyStudyLogs[dateCounter].Date = date.Format("2006-01-02")
		if dailyStudyLogs[dateCounter].Log == nil {
			dailyStudyLogs[dateCounter].Log = []dailySubjectLog{}
		}
		dateCounter++
		date = date.AddDate(0, 0, 1)

		if logCounter == len(studylogs) {
			for dateCounter < termDate {
				dailyStudyLogs[dateCounter].Date = date.Format("2006-01-02")
				dailyStudyLogs[dateCounter].Log = []dailySubjectLog{}
				dateCounter++
				date = date.AddDate(0, 0, 1)
			}
			break
		}
	}

	return c.JSON(dailyStudyLogs)
}
