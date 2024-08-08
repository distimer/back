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
	"pentag.kr/distimer/ent/group"
	"pentag.kr/distimer/ent/studylog"
	"pentag.kr/distimer/middlewares"
	"pentag.kr/distimer/utils/logger"
)

type groupMemberStatisticsElem struct {
	Subject   subjectctrl.SubjectDTO `json:"subject" validate:"required"`
	StudyTime int                    `json:"study_time" validate:"required"`
}

type groupMemberStatisticscResponse struct {
	UserID uuid.UUID                   `json:"user_id" validate:"required"`
	Log    []groupMemberStatisticsElem `json:"log" validate:"required"`
}

// @Summary Get Group Member Statistics by Date
// @Tags StudyLog
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "group id"
// @Param date query string false "2006-01-02"
// @Success 200 {array} groupMemberStatisticscResponse
// @Failure 400
// @Failure 403
// @Failure 404
// @Failure 500
// @Router /studylog/group/statistics/date/{id} [get]
func GroupMemberStatisticsByDate(c *fiber.Ctx) error {

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

	groupIDStr := c.Params("id")
	groupID, err := uuid.Parse(groupIDStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid group id",
		})
	}

	userID := middlewares.GetUserIDFromMiddleware(c)

	dbConn := db.GetDBClient()

	groupObj, err := dbConn.Group.Query().Where(group.ID(groupID)).WithOwner().Only(context.Background())
	if err != nil {
		if ent.IsNotFound(err) {
			return c.Status(404).JSON(fiber.Map{
				"error": "Group not found",
			})
		}
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	if groupObj.Edges.Owner.ID != userID {
		userAffilitaionObj, err := dbConn.Affiliation.Query().Where(
			affiliation.And(
				affiliation.GroupID(groupID),
				affiliation.UserID(userID),
			),
		).Only(context.Background())
		if err != nil {
			if ent.IsNotFound(err) {
				return c.Status(404).JSON(fiber.Map{
					"error": "You are not the member of the group",
				})
			}
			logger.Error(c, err)
			return c.Status(500).JSON(fiber.Map{
				"error": "Internal server error",
			})
		}
		if userAffilitaionObj.Role < groupObj.RevealPolicy {
			return c.Status(403).JSON(fiber.Map{
				"error": "You are not allowed to see the log",
			})
		}
	}
	// get all members
	members, err := dbConn.Affiliation.Query().Where(affiliation.GroupID(groupID)).All(context.Background())
	if err != nil {
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	// get all shared study logs
	studyLogList, err := dbConn.StudyLog.Query().Where(
		studylog.And(
			studylog.StartAtLTE(date.AddDate(0, 0, 1)),
			studylog.EndAtGTE(date),
			studylog.HasSharedGroupWith(group.ID(groupID)),
		),
	).WithUser().WithSubject().All(context.Background())
	if err != nil {
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	result := make([]groupMemberStatisticscResponse, len(members))
	for i, member := range members {
		result[i].UserID = member.UserID
		for _, log := range studyLogList {
			if log.Edges.User.ID == member.UserID {
				if log.StartAt.Before(date) {
					result[i].Log = append(result[i].Log, groupMemberStatisticsElem{
						Subject: subjectctrl.SubjectDTO{
							ID:    log.Edges.Subject.ID.String(),
							Name:  log.Edges.Subject.Name,
							Color: log.Edges.Subject.Color,
							Order: log.Edges.Subject.Order,
						},
						StudyTime: int(log.EndAt.Sub(date).Minutes()),
					})
				} else if log.EndAt.After(date.AddDate(0, 0, 1)) {
					result[i].Log = append(result[i].Log, groupMemberStatisticsElem{
						Subject: subjectctrl.SubjectDTO{
							ID:    log.Edges.Subject.ID.String(),
							Name:  log.Edges.Subject.Name,
							Color: log.Edges.Subject.Color,
							Order: log.Edges.Subject.Order,
						},
						StudyTime: int(date.AddDate(0, 0, 1).Sub(log.StartAt).Minutes()),
					})
				} else {

					result[i].Log = append(result[i].Log, groupMemberStatisticsElem{
						Subject: subjectctrl.SubjectDTO{
							ID:    log.Edges.Subject.ID.String(),
							Name:  log.Edges.Subject.Name,
							Color: log.Edges.Subject.Color,
							Order: log.Edges.Subject.Order,
						},
						StudyTime: int(log.EndAt.Sub(log.StartAt).Minutes()),
					})
				}
			}
		}

	}

	return c.JSON(result)
}
