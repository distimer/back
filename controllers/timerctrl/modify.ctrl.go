package timerctrl

import (
	"context"
	"unicode/utf8"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"pentag.kr/distimer/db"
	"pentag.kr/distimer/ent"
	"pentag.kr/distimer/ent/affiliation"
	"pentag.kr/distimer/ent/subject"
	"pentag.kr/distimer/ent/timer"
	"pentag.kr/distimer/middlewares"
	"pentag.kr/distimer/utils/dto"
	"pentag.kr/distimer/utils/logger"
)

type modifyTimerDTO struct {
	SubjectID      string   `json:"subject_id" validate:"required,uuid"`
	Content        string   `json:"content" validate:"required" example:"content between 0 and 30"`
	SharedGroupIDs []string `json:"shared_group_ids" validate:"required"`
}

// @Summary Modify Timer
// @Tags Timer
// @Accept json
// @Produce json
// @Security Bearer
// @Param request body modifyTimerDTO true "modifyTimerDTO"
// @Success 200 {object} timerDTO
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /timer [put]
func ModifyTimer(c *fiber.Ctx) error {
	data := new(modifyTimerDTO)
	if err := dto.Bind(c, data); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err,
		})
	}
	if utf8.RuneCountInString(data.Content) > 30 {
		return c.Status(400).JSON(fiber.Map{
			"error": "Content length should be between 1 and 30",
		})
	}

	subjectID := uuid.MustParse(data.SubjectID)
	userID := middlewares.GetUserIDFromMiddleware(c)

	dbConn := db.GetDBClient()

	subjectExist, err := dbConn.Subject.Query().Where(subject.ID(subjectID)).Exist(context.Background())
	if err != nil {
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	} else if !subjectExist {
		return c.Status(404).JSON(fiber.Map{
			"error": "Subject not found",
		})
	}

	timerObj, err := dbConn.Timer.Query().Where(timer.UserID(userID)).Only(context.Background())
	if err != nil {
		if ent.IsNotFound(err) {
			return c.Status(404).JSON(fiber.Map{
				"error": "Timer not found",
			})
		}
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
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
			logger.Error(c, err)
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

	timerObj, err = timerObj.Update().
		Where(timer.UserID(userID)).
		SetContent(data.Content).
		SetSubjectID(subjectID).
		ClearSharedGroup().
		AddSharedGroupIDs(sharedGroupIDs...).
		Save(context.Background())
	if err != nil {
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}
	return c.JSON(timerDTO{
		ID:             timerObj.ID.String(),
		SubjectID:      data.SubjectID,
		Content:        timerObj.Content,
		SharedGroupIDs: data.SharedGroupIDs,
	})
}
