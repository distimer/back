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
	"pentag.kr/distimer/ent/user"
	"pentag.kr/distimer/middlewares"
	"pentag.kr/distimer/utils/logger"
)

// @Summary Get StudyLogs with Term with Group
// @Tags StudyLog
// @Accept json
// @Produce json
// @Security Bearer
// @Param group_id path string true "group id"
// @Param member_id path string true "member id"
// @Param start_date query string true "2006-01-02"
// @Param end_date query string true "2006-01-03"
// @Success 200 {array} groupStudyLogDTO
// @Failure 400
// @Failure 403
// @Failure 404
// @Failure 500
// @Router /studylog/group/term/{group_id}/{member_id} [get]
func GetByTermWithGroup(c *fiber.Ctx) error {
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
		logger.CtxError(c, err)
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
		logger.CtxError(c, err)
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
	logList, err := dbConn.StudyLog.Query().
		Where(
			studylog.And(
				studylog.And(
					studylog.HasUserWith(user.ID(memberID)),
					studylog.HasSharedGroupWith(group.ID(groupID)),
				),
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

	result := make([]groupStudyLogDTO, len(logList))
	for i, log := range logList {
		categoryObj, err := log.Edges.Subject.QueryCategory().Only(context.Background())
		if err != nil {
			logger.CtxError(c, err)
			return c.Status(500).JSON(fiber.Map{
				"error": "Internal server error",
			})
		}
		result[i] = groupStudyLogDTO{
			ID: log.ID.String(),
			Subject: subjectctrl.SubjectDTO{
				ID:    log.Edges.Subject.ID.String(),
				Name:  log.Edges.Subject.Name,
				Color: log.Edges.Subject.Color,
				Order: log.Edges.Subject.Order,
			},
			CategoryID:   categoryObj.ID.String(),
			CategoryName: categoryObj.Name,
			StartAt:      log.StartAt.Format(time.RFC3339),
			EndAt:        log.EndAt.Format(time.RFC3339),
			Content:      log.Content,
		}
	}
	return c.JSON(result)
}
