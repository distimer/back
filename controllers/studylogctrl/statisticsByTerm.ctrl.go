package studylogctrl

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"pentag.kr/distimer/db"
	"pentag.kr/distimer/ent"
	"pentag.kr/distimer/ent/studylog"
	"pentag.kr/distimer/utils/logger"
)

const (
	maxTermDateRange = 31
)

type dailySubjectLog struct {
	SubjectID string `json:"subject_id"`
	StudyTime int    `json:"study_time"`
}

type dailyStudyLog struct {
	Date time.Time         `json:"date"`
	Log  []dailySubjectLog `json:"log"`
}

func GetStatisticsWithTerm(c *fiber.Ctx) error {
	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")

	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid start_date",
		})
	}
	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid end_date",
		})
	}
	if startDate.After(endDate) {
		return c.Status(400).JSON(fiber.Map{
			"error": "start_date should be before end_date",
		})
	}
	// term date
	termDate := int(endDate.Sub(startDate).Hours() / 24)
	if termDate > maxTermDateRange {
		return c.Status(400).JSON(fiber.Map{
			"error": "The term date range should be less than 31 days",
		})
	}

	dbConn := db.GetDBClient()

	studylogs, err := dbConn.StudyLog.Query().Where(studylog.StartAtLTE(endDate), studylog.EndAtLTE(startDate)).Order(ent.Asc("start_at")).All(context.Background())
	if err != nil {
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	// Calculate the total study time

	// create array of daily study logs
	dailyStudyLogs := make([]dailyStudyLog, termDate)
	dateCounter := 0
	logCounter := 0
	date := startDate
	for dateCounter < termDate {

		dailyStudyLogs[dateCounter].Date = date
		firstLog := studylogs[logCounter]

		// check if the log is separated by date
		if firstLog.StartAt.Day() < date.Day() {
			dailyStudyLogs[dateCounter].Log = append(
				dailyStudyLogs[dateCounter].Log,
				dailySubjectLog{
					SubjectID: firstLog.Edges.Subject.ID.String(),
					StudyTime: int(firstLog.EndAt.Sub(date).Seconds()),
				},
			)
			logCounter++
		}

		flag := true
		for flag && studylogs[logCounter].StartAt.Day() <= date.Day() {
			log := studylogs[logCounter]
			if log.EndAt.Day() > date.Day() {
				flag = false
				dailyStudyLogs[dateCounter].Log = append(
					dailyStudyLogs[dateCounter].Log,
					dailySubjectLog{
						SubjectID: log.Edges.Subject.ID.String(),
						StudyTime: int(date.Sub(log.StartAt).Seconds()),
					},
				)
			} else {
				dailyStudyLogs[dateCounter].Log = append(
					dailyStudyLogs[dateCounter].Log,
					dailySubjectLog{
						SubjectID: log.Edges.Subject.ID.String(),
						StudyTime: int(log.EndAt.Sub(log.StartAt).Seconds()),
					},
				)
				logCounter++
			}
		}
		dateCounter++
		date = date.AddDate(0, 0, 1)
	}

	return c.JSON(dailyStudyLogs)
}
