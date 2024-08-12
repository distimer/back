package groupctrl

import (
	"context"
	"time"
	"unicode/utf8"

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

type modifyGroupReq struct {
	Name           string `json:"name" validate:"required" example:"name between 3 and 30"`
	Description    string `json:"description" example:"description between 0 and 100"`
	NicknamePolicy string `json:"nickname_policy" example:"nickname_policy between 0 and 50"`
	RevealPolicy   *int8  `json:"reveal_policy" validate:"required,min=0,max=2"`
	InvitePolicy   *int8  `json:"invite_policy" validate:"required,min=0,max=2"`
}

// @Summary Modify Group Info
// @Tags Group
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "group id"
// @Param request body groupctrl.modifyGroupReq true "modifyGroupReq"
// @Success 200 {object} groupDTO
// @Router /group/{id} [put]
func ModifyGroupInfo(c *fiber.Ctx) error {
	groupIDStr := c.Params("id")
	groupID, err := uuid.Parse(groupIDStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid group id",
		})
	}

	data := new(modifyGroupReq)
	if err := dto.Bind(c, data); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err,
		})
	}
	if utf8.RuneCountInString(data.Description) > 100 {
		return c.Status(400).JSON(fiber.Map{
			"error": "Description length should be less than 100",
		})
	} else if utf8.RuneCountInString(data.NicknamePolicy) > 50 {
		return c.Status(400).JSON(fiber.Map{
			"error": "Nickname policy length should be less than 50",
		})
	} else if utf8.RuneCountInString(data.Name) < 3 || utf8.RuneCountInString(data.Name) > 30 {
		return c.Status(400).JSON(fiber.Map{
			"error": "Name length should be less than 30",
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

	err = groupObj.Update().
		SetName(data.Name).
		SetDescription(data.Description).
		SetNicknamePolicy(data.NicknamePolicy).
		SetRevealPolicy(*data.RevealPolicy).
		SetInvitePolicy(*data.InvitePolicy).
		Exec(context.Background())

	if err != nil {
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	affiliationObj, err := dbConn.Affiliation.Query().Where(affiliation.And(affiliation.GroupID(groupID), affiliation.UserID(userID))).Only(context.Background())

	if err != nil {
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}
	return c.JSON(groupDTO{
		ID:             groupID.String(),
		Name:           data.Name,
		Description:    data.Description,
		NicknamePolicy: data.NicknamePolicy,
		RevealPolicy:   *data.RevealPolicy,
		InvitePolicy:   *data.InvitePolicy,
		CreateAt:       groupObj.CreatedAt.Format(time.RFC3339),
		OwnerNickname:  affiliationObj.Nickname,
	})
}
