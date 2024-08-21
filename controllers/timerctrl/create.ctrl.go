package timerctrl

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"pentag.kr/distimer/db"
	"pentag.kr/distimer/ent/affiliation"
	"pentag.kr/distimer/ent/subject"
	"pentag.kr/distimer/ent/timer"
	"pentag.kr/distimer/middlewares"
	"pentag.kr/distimer/utils/dto"
	"pentag.kr/distimer/utils/logger"
)

func CreateTimer(c *fiber.Ctx) error {
	data := new(timerMetadataDTO)
	if err := dto.Bind(c, data); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err,
		})
	}

	subjectID := uuid.MustParse(data.SubjectID)

	userID := middlewares.GetUserIDFromMiddleware(c)

	dbConn := db.GetDBClient()

	subjectExist, err := dbConn.Subject.Query().Where(subject.ID(subjectID)).Exist(context.Background())
	if err != nil {
		logger.CtxError(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	} else if !subjectExist {
		return c.Status(404).JSON(fiber.Map{
			"error": "Subject not found",
		})
	}

	timerExist, err := dbConn.Timer.Query().Where(timer.UserID(userID)).Exist(context.Background())
	if err != nil {
		logger.CtxError(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	} else if timerExist {
		return c.Status(409).JSON(fiber.Map{
			"error": "Timer already exists",
		})
	}

	var sharedGroupIDs []uuid.UUID
	for _, groupIDStr := range data.SharedGroupIDs {
		groupID, err := uuid.Parse(groupIDStr)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Invalid group ID",
			})
		}

		exist, err := dbConn.Affiliation.Query().Where(affiliation.And(affiliation.GroupID(groupID), affiliation.UserID(userID))).Exist(context.Background())
		if err != nil {
			logger.CtxError(c, err)
			return c.Status(500).JSON(fiber.Map{
				"error": "Internal server error",
			})
		}
		if !exist {
			return c.Status(403).JSON(fiber.Map{
				"error": "You are not the member of the group",
			})
		}
		sharedGroupIDs = append(sharedGroupIDs, groupID)
	}

	timer, err := dbConn.Timer.Create().
		SetContent(data.Content).
		SetSubjectID(subjectID).
		SetUserID(userID).
		AddSharedGroupIDs(sharedGroupIDs...).
		Save(context.Background())
	if err != nil {
		logger.CtxError(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}
	return c.JSON(
		timerDTO{
			ID:             timer.ID.String(),
			SubjectID:      subjectID.String(),
			Content:        timer.Content,
			StartAt:        timer.StartAt.Format(time.RFC3339),
			SharedGroupIDs: data.SharedGroupIDs,
		},
	)
}
