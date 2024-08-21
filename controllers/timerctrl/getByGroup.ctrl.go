package timerctrl

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"pentag.kr/distimer/controllers/groupctrl"
	"pentag.kr/distimer/controllers/subjectctrl"
	"pentag.kr/distimer/db"
	"pentag.kr/distimer/ent"
	"pentag.kr/distimer/ent/affiliation"
	"pentag.kr/distimer/middlewares"
	"pentag.kr/distimer/utils/logger"
)

func GetTimerByGroup(c *fiber.Ctx) error {

	groupIDStr := c.Params("id")
	groupID, err := uuid.Parse(groupIDStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid group ID",
		})
	}

	userID := middlewares.GetUserIDFromMiddleware(c)

	dbConn := db.GetDBClient()

	affiliationObj, err := dbConn.Affiliation.Query().
		Where(affiliation.GroupID(groupID), affiliation.UserID(userID)).
		WithGroup().
		Only(context.Background())
	if err != nil {
		if ent.IsNotFound(err) {
			return c.Status(404).JSON(fiber.Map{
				"error": "Group not found or not affiliated",
			})
		}
		logger.CtxError(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}
	if affiliationObj.Edges.Group.RevealPolicy > affiliationObj.Role {
		return c.Status(403).JSON(fiber.Map{
			"error": "Not Enough Permission",
		})
	}

	timers, err := affiliationObj.Edges.Group.QuerySharedTimer().WithUser().WithSubject().All(context.Background())
	if err != nil {
		logger.CtxError(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}
	result := make([]timerWithEdgeInfoDTO, len(timers))
	for i, timer := range timers {

		affiliationObj, err := timer.Edges.User.QueryAffiliations().Where(affiliation.GroupID(groupID)).Only(context.Background())
		if err != nil {
			logger.CtxError(c, err)
			return c.Status(500).JSON(fiber.Map{
				"error": "Internal server error",
			})
		}
		result[i] = timerWithEdgeInfoDTO{
			ID: timer.ID.String(),
			Subject: subjectctrl.SubjectDTO{
				ID:    timer.Edges.Subject.ID.String(),
				Name:  timer.Edges.Subject.Name,
				Color: timer.Edges.Subject.Color,
			},
			Content: timer.Content,
			StartAt: timer.StartAt.Format(time.RFC3339),
			Affiliation: groupctrl.AffiliationDTO{
				GroupID:  affiliationObj.GroupID.String(),
				UserID:   affiliationObj.UserID.String(),
				Role:     affiliationObj.Role,
				Nickname: affiliationObj.Nickname,
				JoinedAt: affiliationObj.JoinedAt.Format(time.RFC3339),
			},
		}
	}
	return c.JSON(result)
}
