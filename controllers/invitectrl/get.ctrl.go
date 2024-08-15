package invitectrl

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"pentag.kr/distimer/configs"
	"pentag.kr/distimer/db"
	"pentag.kr/distimer/ent"
	"pentag.kr/distimer/ent/affiliation"
	"pentag.kr/distimer/ent/group"
	"pentag.kr/distimer/ent/invitecode"
	"pentag.kr/distimer/utils/logger"
)

type inviteCodeInfoRes struct {
	GroupName          string `json:"group_name" validate:"required"`
	GroupOwnerNickname string `json:"group_owner_nickname" validate:"required"`
	GroupDescription   string `json:"group_description" validate:"required"`
	NicknamePolicy     string `json:"nickname_policy" validate:"required"`
}

// @Summary Get Invite Code Info
// @Tags Invite
// @Accept json
// @Produce json
// @Security Bearer
// @Param code path string true "invite code"
// @Success 200 {object} inviteCodeInfoRes
// Failure 400
// Failure 404
// Failure 500
// @Router /invite/{code} [get]
func GetInviteCodeInfo(c *fiber.Ctx) error {
	code := c.Params("code")
	if len(code) != configs.InviteCodeLength {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid invite code",
		})
	}

	dbConn := db.GetDBClient()

	groupObj, err := dbConn.Group.Query().
		Where(group.HasInviteCodesWith(invitecode.Code(code))).
		WithOwner().
		Only(context.Background())
	if err != nil {
		if ent.IsNotFound(err) {
			return c.Status(404).JSON(fiber.Map{
				"error": "Invite code is not exist",
			})
		}
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	ownerAffiliation, err := dbConn.Affiliation.Query().
		Where(affiliation.And(affiliation.GroupID(groupObj.ID), affiliation.Role(2))).
		Only(context.Background())
	if err != nil {
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}
	return c.JSON(inviteCodeInfoRes{
		GroupName:          groupObj.Name,
		GroupOwnerNickname: ownerAffiliation.Nickname,
		GroupDescription:   groupObj.Description,
		NicknamePolicy:     groupObj.NicknamePolicy,
	})
}
