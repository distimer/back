package groupctrl

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"pentag.kr/distimer/db"
	"pentag.kr/distimer/ent"
	"pentag.kr/distimer/ent/group"
	"pentag.kr/distimer/ent/user"
	"pentag.kr/distimer/middlewares"
	"pentag.kr/distimer/utils/dto"
	"pentag.kr/distimer/utils/logger"
)

type modifyGroupPolicyReq struct {
	RevealPolicy int8 `json:"reveal_policy" validate:"required,min=0,max=2"`
	InvitePolicy int8 `json:"invite_policy" validate:"required,min=0,max=2"`
}
type modifyGroupPolicyRes struct {
	RevealPolicy int8 `json:"reveal_policy" validate:"required,min=0,max=2"`
	InvitePolicy int8 `json:"invite_policy" validate:"required,min=0,max=2"`
}

// @Summary Modify Group Policy
// @Tags Group
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "group id"
// @Param request body groupctrl.modifyGroupPolicyReq true "modifyGroupPolicyReq"
// @Success 200 {object} modifyGroupPolicyRes
// @Router /group/policy/{id} [put]
func ModifyGroupPolicy(c *fiber.Ctx) error {
	groupIDStr := c.Params("id")
	groupID, err := uuid.Parse(groupIDStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid group id",
		})
	}

	data := new(modifyGroupPolicyReq)
	if err := dto.Bind(c, data); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err,
		})
	}

	userID := middlewares.GetUserIDFromMiddleware(c)
	dbConn := db.GetDBClient()

	groupObj, err := dbConn.Group.Query().Where(group.And(group.ID(groupID), group.HasOwnerWith(user.ID(userID)))).Only(context.Background())
	if err != nil {
		if ent.IsNotFound(err) {
			return c.Status(404).JSON(fiber.Map{
				"error": "Group not found or you are not the owner",
			})
		}
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	err = groupObj.Update().SetRevealPolicy(data.RevealPolicy).SetInvitePolicy(data.InvitePolicy).Exec(context.Background())
	if err != nil {
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	return c.Status(200).JSON(modifyGroupPolicyRes{
		RevealPolicy: data.RevealPolicy,
		InvitePolicy: data.InvitePolicy,
	})
}
