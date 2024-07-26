package groupctrl

import (
	"context"
	"math/rand"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"pentag.kr/distimer/db"
	"pentag.kr/distimer/ent"
	"pentag.kr/distimer/ent/affiliation"
	"pentag.kr/distimer/ent/group"
	"pentag.kr/distimer/ent/invitecode"
	"pentag.kr/distimer/middlewares"
	"pentag.kr/distimer/utils/logger"
)

const InviteCodePerGroupLimit = 3

type inviteGroupRes struct {
	Code string `json:"code"`
}

// @Summary Invite to Group
// @Tags Group
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "group id"
// @Success 200 {object} inviteGroupRes
// @Router /group/invite/{id} [post]
func InviteToGroup(c *fiber.Ctx) error {
	groupIDStr := c.Params("id")
	groupID, err := uuid.Parse(groupIDStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid group id",
		})
	}

	userID := middlewares.GetUserIDFromMiddleware(c)

	dbConn := db.GetDBClient()

	affiliationObj, err := dbConn.Affiliation.Query().Where(affiliation.And(affiliation.GroupID(groupID), affiliation.UserID(userID))).Only(context.Background())
	if err != nil {
		if ent.IsNotFound(err) {
			return c.Status(404).JSON(fiber.Map{
				"error": "Group is not exist , or you are not the member of the group",
			})
		}
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}
	if affiliationObj.Role != 2 {
		groupObj, err := dbConn.Group.Query().Where(group.ID(groupID)).Only(context.Background())
		if err != nil {
			logger.Error(c, err)
			return c.Status(500).JSON(fiber.Map{
				"error": "Internal server error",
			})
		} else if affiliationObj.Role < groupObj.InvitePolicy {
			return c.Status(403).JSON(fiber.Map{
				"error": "You are not allowed to invite",
			})
		}
	}
	count, err := dbConn.InviteCode.Query().Where(invitecode.HasGroupWith(group.ID(groupID))).Count(context.Background())
	if err != nil {
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	} else if count >= InviteCodePerGroupLimit {
		return c.Status(409).JSON(fiber.Map{
			"error": "Invite code limit exceeded",
		})
	}

	newInviteCode := randString(7)

	_, err = dbConn.InviteCode.Create().
		SetCode(newInviteCode).
		SetGroupID(groupID).
		Save(context.Background())
	if err != nil {
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}
	return c.JSON(inviteGroupRes{
		Code: newInviteCode,
	})
}

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func randString(n int) string {
	// 문자열의 길이가 0인 경우 빈 문자열 반환
	if n <= 0 {
		return ""
	}

	// 랜덤한 문자를 저장할 바이트 슬라이스 생성
	b := make([]byte, n)
	for i := range b {
		// letters 문자열에서 랜덤한 문자 선택
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// @Summary Get Invite Code List
// @Tags Group
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "group id"
// @Success 200 {array} string
// @Router /group/invite/{id} [get]
func GetInviteCodeList(c *fiber.Ctx) error {
	groupIDStr := c.Params("id")
	groupID, err := uuid.Parse(groupIDStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid group id",
		})
	}

	userID := middlewares.GetUserIDFromMiddleware(c)

	dbConn := db.GetDBClient()

	affiliationObj, err := dbConn.Affiliation.Query().Where(affiliation.And(affiliation.GroupID(groupID), affiliation.UserID(userID))).Only(context.Background())
	if err != nil {
		if ent.IsNotFound(err) {
			return c.Status(404).JSON(fiber.Map{
				"error": "Group is not exist , or you are not the member of the group",
			})
		}
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}
	if affiliationObj.Role != 2 {
		groupObj, err := dbConn.Group.Query().Where(group.ID(groupID)).Only(context.Background())
		if err != nil {
			logger.Error(c, err)
			return c.Status(500).JSON(fiber.Map{
				"error": "Internal server error",
			})
		} else if affiliationObj.Role < groupObj.InvitePolicy {
			return c.Status(403).JSON(fiber.Map{
				"error": "You are not allowed to invite",
			})
		}
	}
	inviteCodeList := []string{}
	inviteCodeList, err = dbConn.InviteCode.Query().
		Where(invitecode.HasGroupWith(group.ID(groupID))).
		Select(invitecode.FieldCode).
		Strings(context.Background())
	if err != nil {
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}
	return c.JSON(inviteCodeList)
}

// @Summary Delete Invite Code
// @Tags Group
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "group id"
// @Param code path string true "invite code"
// @Success 204
// @Router /group/invite/{id}/{code} [delete]
func DeleteInviteCode(c *fiber.Ctx) error {
	groupIDStr := c.Params("id")
	groupID, err := uuid.Parse(groupIDStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid group id",
		})
	}

	userID := middlewares.GetUserIDFromMiddleware(c)

	dbConn := db.GetDBClient()

	affiliationObj, err := dbConn.Affiliation.Query().Where(affiliation.And(affiliation.GroupID(groupID), affiliation.UserID(userID))).Only(context.Background())
	if err != nil {
		if ent.IsNotFound(err) {
			return c.Status(404).JSON(fiber.Map{
				"error": "Group is not exist , or you are not the member of the group",
			})
		}
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}
	if affiliationObj.Role != 2 {
		groupObj, err := dbConn.Group.Query().Where(group.ID(groupID)).Only(context.Background())
		if err != nil {
			logger.Error(c, err)
			return c.Status(500).JSON(fiber.Map{
				"error": "Internal server error",
			})
		} else if affiliationObj.Role < groupObj.InvitePolicy {
			return c.Status(403).JSON(fiber.Map{
				"error": "You are not allowed to invite",
			})
		}
	}

	code := c.Params("code")

	_, err = dbConn.InviteCode.Delete().
		Where(invitecode.And(invitecode.HasGroupWith(group.ID(groupID)), invitecode.Code(code))).
		Exec(context.Background())
	if err != nil {
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}
	return c.SendStatus(204)
}
