package groupctrl

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"pentag.kr/distimer/db"
	"pentag.kr/distimer/ent/user"
	"pentag.kr/distimer/middlewares"
	"pentag.kr/distimer/utils/logger"
)

// @Summary Get All Joined Groups
// @Tags Group
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {array} groupDTO
// @Router /group [get]
func GetJoinedGroups(c *fiber.Ctx) error {
	userID := middlewares.GetUserIDFromMiddleware(c)

	dbConn := db.GetDBClient()
	groups, err := dbConn.User.Query().Where(user.ID(userID)).QueryJoinedGroups().All(context.Background())
	if err != nil {
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	result := make([]groupDTO, len(groups))
	for i, group := range groups {
		result[i] = groupDTO{
			ID:             group.ID.String(),
			Name:           group.Name,
			Description:    group.Description,
			NicknamePolicy: group.NicknamePolicy,
			RevealPolicy:   group.RevealPolicy,
			InvitePolicy:   group.InvitePolicy,
			CreateAt:       group.CreatedAt.Format(time.RFC3339),
		}
	}

	return c.JSON(result)

}
