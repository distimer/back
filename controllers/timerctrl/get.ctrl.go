package timerctrl

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"pentag.kr/distimer/db"
	"pentag.kr/distimer/ent"
	"pentag.kr/distimer/ent/timer"
	"pentag.kr/distimer/ent/user"
	"pentag.kr/distimer/middlewares"
	"pentag.kr/distimer/utils/logger"
)

// @Summary Get My Timer Info
// @Tags Timer
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} timerDTO
// @Success 404
// @Failure 500
// @Router /timer [get]
func GetMyTimerInfo(c *fiber.Ctx) error {
	userID := middlewares.GetUserIDFromMiddleware(c)

	dbConn := db.GetDBClient()

	foundTimer, err := dbConn.Timer.Query().Where(timer.HasUserWith(user.ID(userID))).WithSharedGroup().First(context.Background())
	if err != nil {
		if ent.IsNotFound(err) {
			return c.Status(404).JSON(fiber.Map{
				"info": "Timer not found",
			})
		}
		logger.CtxError(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}
	return c.JSON(timerDTO{
		ID:        foundTimer.ID.String(),
		SubjectID: foundTimer.SubjectID.String(),
		Content:   foundTimer.Content,
		StartAt:   foundTimer.StartAt.Format(time.RFC3339),
		SharedGroupIDs: func() []string {
			sharedGroups := foundTimer.Edges.SharedGroup
			result := make([]string, len(sharedGroups))
			for i, sharedGroup := range sharedGroups {
				result[i] = sharedGroup.ID.String()
			}
			return result
		}(),
	})
}
