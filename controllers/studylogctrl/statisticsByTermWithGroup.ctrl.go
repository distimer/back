package studylogctrl

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"pentag.kr/distimer/controllers/subjectctrl"
	"pentag.kr/distimer/db"
	"pentag.kr/distimer/ent"
	"pentag.kr/distimer/ent/affiliation"
	"pentag.kr/distimer/ent/studylog"
	"pentag.kr/distimer/ent/user"
	"pentag.kr/distimer/middlewares"
	"pentag.kr/distimer/utils/logger"
)

type groupMemberdailySubjectLog struct {
	Subject      subjectctrl.SubjectDTO `json:"subject_id" validate:"required"`
	CategoryID   string                 `json:"category_id" validate:"required"`
	CategoryName string                 `json:"category_name" validate:"required"`
	StudyTime    int                    `json:"study_time" validate:"required"`
}

type groupMemberdailyStudyLog struct {
	Date string                       `json:"date" validate:"required"`
	Log  []groupMemberdailySubjectLog `json:"log" validate:"required"`
}

// @Summary Get Statistics with Term
// @Tags StudyLog
// @Accept json
// @Produce json
// @Security Bearer
// @Param start_date query string true "2006-01-02"
// @Param end_date query string true "2006-01-03"
// @Param group_id path string true "group_id"
// @Param member_id path string true "member_id"
// @Success 200 {array} groupMemberdailyStudyLog
// @Failure 400
// @Failure 403
// @Failure 404
// @Failure 500
// @Router /studylog/statistics/term/{group_id}/{member_id} [get]
func GetStatisticsByTermWithGroup(c *fiber.Ctx) error {
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

	groupIDStr := c.Params("group_id")
	memberIDStr := c.Params("member_id")

	groupID, err := uuid.Parse(groupIDStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid group ID",
		})
	}
	memberID, err := uuid.Parse(memberIDStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid member ID",
		})
	}

	userID := middlewares.GetUserIDFromMiddleware(c)
	dbConn := db.GetDBClient()

	userAffiliationObj, err := dbConn.Affiliation.Query().
		Where(
			affiliation.And(
				affiliation.GroupID(groupID),
				affiliation.UserID(userID),
			),
		).
		WithGroup().
		Only(context.Background())
	if err != nil {
		if ent.IsNotFound(err) {
			return c.Status(404).JSON(fiber.Map{
				"error": "You are not a member of this group",
			})
		}
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	if userAffiliationObj.Role < userAffiliationObj.Edges.Group.RevealPolicy {
		return c.Status(403).JSON(fiber.Map{
			"error": "You don't have permission to view study logs",
		})
	}

	memberAffiliationExist, err := dbConn.Affiliation.Query().
		Where(
			affiliation.And(
				affiliation.GroupID(groupID),
				affiliation.UserID(memberID),
			),
		).
		Exist(context.Background())
	if err != nil {
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}
	if !memberAffiliationExist {
		return c.Status(404).JSON(fiber.Map{
			"error": "Member not found",
		})
	}

	// get studylog
	studylogs, err := dbConn.StudyLog.Query().
		Where(
			studylog.And(
				studylog.HasUserWith(user.ID(memberID)),
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
		return c.JSON([]groupMemberdailyStudyLog{})
	}

	// create array of daily study logs
	dailyStudyLogs := make([]groupMemberdailyStudyLog, termDate)
	dateCounter := 0
	logCounter := 0
	date := startDate
	for dateCounter < termDate {

		firstLog := studylogs[logCounter]

		// check if the log is separated by date
		if firstLog.StartAt.Day() < date.Day() {
			dailyStudyLogs[dateCounter].Log = append(
				dailyStudyLogs[dateCounter].Log,
				groupMemberdailySubjectLog{
					Subject: subjectctrl.SubjectDTO{
						ID:    firstLog.Edges.Subject.ID.String(),
						Name:  firstLog.Edges.Subject.Name,
						Color: firstLog.Edges.Subject.Color,
						Order: firstLog.Edges.Subject.Order,
					},
					CategoryID:   firstLog.Edges.Subject.Edges.Category.ID.String(),
					CategoryName: firstLog.Edges.Subject.Edges.Category.Name,
					StudyTime:    int(firstLog.EndAt.Sub(date).Seconds()),
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
					groupMemberdailySubjectLog{
						Subject: subjectctrl.SubjectDTO{
							ID:    log.Edges.Subject.ID.String(),
							Name:  log.Edges.Subject.Name,
							Color: log.Edges.Subject.Color,
							Order: log.Edges.Subject.Order,
						},
						CategoryID:   log.Edges.Subject.Edges.Category.ID.String(),
						CategoryName: log.Edges.Subject.Edges.Category.Name,
						StudyTime:    int(date.AddDate(0, 0, 1).Sub(log.StartAt).Seconds()),
					},
				)
				break
			} else {
				dailyStudyLogs[dateCounter].Log = append(
					dailyStudyLogs[dateCounter].Log,
					groupMemberdailySubjectLog{
						Subject: subjectctrl.SubjectDTO{
							ID:    log.Edges.Subject.ID.String(),
							Name:  log.Edges.Subject.Name,
							Color: log.Edges.Subject.Color,
							Order: log.Edges.Subject.Order,
						},
						CategoryID:   log.Edges.Subject.Edges.Category.ID.String(),
						CategoryName: log.Edges.Subject.Edges.Category.Name,
						StudyTime:    int(log.EndAt.Sub(log.StartAt).Seconds()),
					},
				)
				logCounter++
			}
		}
		dailyStudyLogs[dateCounter].Date = date.Format("2006-01-02")
		if dailyStudyLogs[dateCounter].Log == nil {
			dailyStudyLogs[dateCounter].Log = []groupMemberdailySubjectLog{}
		}
		dateCounter++
		date = date.AddDate(0, 0, 1)

		if logCounter == len(studylogs) {
			for dateCounter < termDate {
				dailyStudyLogs[dateCounter].Date = date.Format("2006-01-02")
				dailyStudyLogs[dateCounter].Log = []groupMemberdailySubjectLog{}
				dateCounter++
				date = date.AddDate(0, 0, 1)
			}
			break
		}
	}
	return nil
}
