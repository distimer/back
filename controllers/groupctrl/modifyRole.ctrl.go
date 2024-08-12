package groupctrl

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"pentag.kr/distimer/db"
	"pentag.kr/distimer/ent"
	"pentag.kr/distimer/ent/affiliation"
	"pentag.kr/distimer/ent/group"
	"pentag.kr/distimer/ent/user"
	"pentag.kr/distimer/middlewares"
	"pentag.kr/distimer/utils/dto"
	"pentag.kr/distimer/utils/logger"
)

type modifyRoleReq struct {
	UserID string `json:"user_id" validate:"required,uuid"`
	Role   *int8  `json:"role" validate:"required,min=0,max=1"`
}

// @Summary Modify Role
// @Tags Group
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "group id"
// @Param request body groupctrl.modifyRoleReq true "modifyRoleReq"
// @Success 200 {object} AffiliationDTO
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /group/role/{id} [patch]
func ModifyRole(c *fiber.Ctx) error {
	groupIDStr := c.Params("id")
	groupID, err := uuid.Parse(groupIDStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid group id",
		})
	}
	data := new(modifyRoleReq)
	if err := dto.Bind(c, data); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err,
		})
	}
	userID := middlewares.GetUserIDFromMiddleware(c)
	targetUserID := uuid.MustParse(data.UserID)
	if userID == targetUserID {
		return c.Status(400).JSON(fiber.Map{
			"error": "You can't modify your role",
		})
	}
	dbConn := db.GetDBClient()

	exist, err := dbConn.Group.Query().Where(group.And(group.ID(groupID), group.HasOwnerWith(user.ID(userID)))).Exist(context.Background())
	if err != nil {
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	} else if !exist {
		return c.Status(404).JSON(fiber.Map{
			"error": "Group not found or you are not the owner",
		})
	}

	affiliationObj, err := dbConn.Affiliation.Query().Where(affiliation.And(affiliation.GroupID(groupID), affiliation.UserID(targetUserID))).Only(context.Background())
	if err != nil {
		if ent.IsNotFound(err) {
			return c.Status(404).JSON(fiber.Map{
				"error": "User not found in the group",
			})
		}
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	err = dbConn.Affiliation.UpdateOne(affiliationObj).SetRole(*data.Role).Exec(context.Background())
	if err != nil {
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	return c.JSON(AffiliationDTO{
		GroupID:  groupID.String(),
		UserID:   data.UserID,
		Nickname: affiliationObj.Nickname,
		Role:     *data.Role,
		JoinedAt: affiliationObj.JoinedAt.Format(time.RFC3339),
	})

}
