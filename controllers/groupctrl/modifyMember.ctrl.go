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

type modifyMemberReq struct {
	Role     *int8  `json:"role" validate:"required,min=0,max=1"`
	Nickname string `json:"nickname" validate:"required" example:"nickname between 1 and 20"`
}

// @Summary Modify Member
// @Tags Group
// @Accept json
// @Produce json
// @Security Bearer
// @Param groupID path string true "group id"
// @Param memberID path string true "member id"
// @Param request body modifyMemberReq true "modifyMemberReq"
// @Success 200 {object} AffiliationDTO
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /group/member/{groupID}/{memberID} [put]
func ModifyMember(c *fiber.Ctx) error {

	groupIDStr := c.Params("group_id")
	memberIDStr := c.Params("member_id")

	groupID, err := uuid.Parse(groupIDStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid group id",
		})
	}
	memberID, err := uuid.Parse(memberIDStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid member id",
		})
	}

	data := new(modifyMemberReq)
	if err := dto.Bind(c, data); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err,
		})
	}

	userID := middlewares.GetUserIDFromMiddleware(c)
	if userID == memberID {
		return c.Status(400).JSON(fiber.Map{
			"error": "You cannot modify yourself",
		})
	}

	dbConn := db.GetDBClient()

	exist, err := dbConn.Group.Query().Where(group.And(group.ID(groupID), group.HasOwnerWith(user.ID(userID)))).Exist(context.Background())
	if err != nil {
		logger.CtxError(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	} else if !exist {
		return c.Status(404).JSON(fiber.Map{
			"error": "Group not found or you are not the owner",
		})
	}

	affiliationObj, err := dbConn.Affiliation.Query().Where(affiliation.And(affiliation.GroupID(groupID), affiliation.UserID(memberID))).Only(context.Background())
	if err != nil {
		if ent.IsNotFound(err) {
			return c.Status(404).JSON(fiber.Map{
				"error": "User not found in the group",
			})
		}
		logger.CtxError(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	err = dbConn.Affiliation.UpdateOne(affiliationObj).SetRole(*data.Role).SetNickname(data.Nickname).Exec(context.Background())
	if err != nil {
		logger.CtxError(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	return c.JSON(AffiliationDTO{
		GroupID:  groupID.String(),
		UserID:   memberIDStr,
		Nickname: data.Nickname,
		Role:     *data.Role,
		JoinedAt: affiliationObj.JoinedAt.Format(time.RFC3339),
	})
}
