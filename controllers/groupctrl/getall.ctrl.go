package groupctrl

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"pentag.kr/distimer/db"
	"pentag.kr/distimer/ent"
	"pentag.kr/distimer/ent/user"
	"pentag.kr/distimer/middlewares"
	"pentag.kr/distimer/utils/logger"
)

type getJoinedGroupsRes struct {
	Groups []*ent.Group `json:"joined_groups"`
}

// @Summary Get All Joined Groups
// @Tags Group
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} getJoinedGroupsRes
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

	return c.JSON(getJoinedGroupsRes{
		Groups: groups,
	})

}
