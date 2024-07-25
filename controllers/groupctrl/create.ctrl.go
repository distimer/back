package groupctrl

import (
	"context"
	"unicode/utf8"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"pentag.kr/distimer/db"
	"pentag.kr/distimer/ent/group"
	"pentag.kr/distimer/ent/user"
	"pentag.kr/distimer/middlewares"
	"pentag.kr/distimer/utils/dto"
	"pentag.kr/distimer/utils/logger"
)

const (
	GroupPerUsreLimit = 5
)

type createGroupReq struct {
	dto.BaseDTO
	Name           string `json:"name" validate:"required" example:"name between 3 and 30"`
	Nickname       string `json:"nickname" validate:"required" example:"nickname between 1 and 20"`
	Description    string `json:"description" example:"description between 0 and 100"`
	NicknamePolicy string `json:"nickname_policy" example:"nickname_policy between 0 and 50"`
	RevealPolicy   int8   `json:"reveal_policy" validate:"required,min=0,max=2"`
	InvitePolicy   int8   `json:"invite_policy" validate:"required,min=0,max=2"`
}

type createGroupRes struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	NicknamePolicy string `json:"nickname_policy"`
	RevealPolicy   int8   `json:"reveal_policy"`
	InvitePolicy   int8   `json:"invite_policy"`
	CreateAt       string `json:"create_at"`
}

// @Summary Create Group
// @Tags Group
// @Accept json
// @Produce json
// @Security Bearer
// @Param request body groupctrl.createGroupReq true "createGroupReq"
// @Success 201
// @Router /group [post]
func CreateGroup(c *fiber.Ctx) error {
	data := new(createGroupReq)
	if err := dto.Bind(c, data); err != nil {
		return err
	}
	if utf8.RuneCountInString(data.Nickname) < 1 || utf8.RuneCountInString(data.Nickname) > 20 {
		return c.Status(400).JSON(fiber.Map{
			"error": "Nickname length should be between 1 and 20",
		})
	} else if utf8.RuneCountInString(data.Description) > 100 {
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

	newGroupID := uuid.New()

	dbConn := db.GetDBClient()
	count, err := dbConn.Group.Query().Where(group.HasOwnerWith(user.ID(userID))).Count(context.Background())
	if err != nil {
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}
	if count >= GroupPerUsreLimit {
		return c.Status(409).JSON(fiber.Map{
			"error": "Group limit exceeded",
		})
	}

	_, err = dbConn.Group.Create().
		SetID(newGroupID).
		SetName(data.Name).
		SetDescription(data.Description).
		SetNicknamePolicy(data.NicknamePolicy).
		SetRevealPolicy(data.RevealPolicy).
		SetInvitePolicy(data.InvitePolicy).
		SetOwnerID(userID).
		Save(context.Background())
	if err != nil {
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	_, err = dbConn.Affiliation.Create().
		SetGroupID(newGroupID).
		SetUserID(userID).
		SetRole(2).
		SetNickname(data.Nickname).
		Save(context.Background())
	if err != nil {
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	return c.Status(201).JSON(createGroupRes{
		ID:             newGroupID.String(),
		Name:           data.Name,
		Description:    data.Description,
		NicknamePolicy: data.NicknamePolicy,
		RevealPolicy:   data.RevealPolicy,
		InvitePolicy:   data.InvitePolicy,
	})
}
