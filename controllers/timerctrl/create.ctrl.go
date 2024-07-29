package timerctrl

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"pentag.kr/distimer/db"
	"pentag.kr/distimer/ent/subject"
	"pentag.kr/distimer/ent/timer"
	"pentag.kr/distimer/middlewares"
	"pentag.kr/distimer/utils/dto"
	"pentag.kr/distimer/utils/logger"
)

type createTimerReq struct {
	SubjectID string `json:"subject_id" validate:"required"`
	Content   string `json:"content" validate:"required"`
}

type createTimerRes struct {
	ID        string `json:"id"`
	SubjectID string `json:"subject_id"`
	Content   string `json:"content"`
	StartAt   string `json:"start_at"`
}

// @Summary Create Timer
// @Tags Timer
// @Accept json
// @Produce json
// @Security Bearer
// @Param request body createTimerReq true "createTimerReq"
// @Success 200 {object} createTimerRes
// @Failure 400
// @Failure 404
// @Failure 409
// @Failure 500
// @Router /timer [post]
func CreateTimer(c *fiber.Ctx) error {
	data := new(createTimerReq)
	if err := dto.Bind(c, data); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err,
		})
	}

	subjectID, err := uuid.Parse(data.SubjectID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid subject ID",
		})
	}

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

	timerExist, err := dbConn.Timer.Query().Where(timer.UserID(userID)).Exist(context.Background())
	if err != nil {
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	} else if timerExist {
		return c.Status(409).JSON(fiber.Map{
			"error": "Timer already exists",
		})
	}

	timer, err := dbConn.Timer.Create().
		SetContent(data.Content).
		SetSubjectID(subjectID).
		SetUserID(userID).
		Save(context.Background())
	if err != nil {
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	return c.JSON(
		createTimerRes{
			ID:        timer.ID.String(),
			SubjectID: subjectID.String(),
			Content:   timer.Content,
			StartAt:   timer.StartAt.Format(time.RFC3339),
		},
	)
}
