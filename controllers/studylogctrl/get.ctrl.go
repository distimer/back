package studylogctrl

import (
	"context"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"pentag.kr/distimer/db"
	"pentag.kr/distimer/ent"
	"pentag.kr/distimer/ent/studylog"
	"pentag.kr/distimer/ent/user"
	"pentag.kr/distimer/middlewares"
)

const (
	maxCount = 20
)

type getStudyLogListRes struct {
	StudyLogs []*ent.StudyLog `json:"studyLogs"`
}

// @Summary Get All My Study Logs
// @Tags StudyLog
// @Accept json
// @Produce json
// @Security Bearer
// @Param count query int true "count"
// @Param offset query int true "offset"
// @Success 200 {object} getStudyLogListRes
// @Router /studylog [get]
func GetAllMyStudyLogs(c *fiber.Ctx) error {
	userID := middlewares.GetUserIDFromMiddleware(c)
	countStr := c.Query("count")
	offsetStr := c.Query("offset")

	count, err := strconv.Atoi(countStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid count",
		})
	} else if count > maxCount || count < 1 {
		return c.Status(400).JSON(fiber.Map{
			"error": "count should be between 1 and 20",
		})
	}
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid offset",
		})
	} else if offset < 0 {
		return c.Status(400).JSON(fiber.Map{
			"error": "offset should be greater than or equal to 0",
		})
	}

	// Fetch study logs from database
	dbConn := db.GetDBClient()
	studyLogs, err := dbConn.StudyLog.Query().
		Where(studylog.HasUserWith(user.ID(userID))).
		Order(ent.Desc("created_at")).
		Limit(count).
		Offset(offset).
		All(context.Background())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}
	return c.JSON(getStudyLogListRes{
		StudyLogs: studyLogs,
	})
}
