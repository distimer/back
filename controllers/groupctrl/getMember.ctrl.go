package groupctrl

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"pentag.kr/distimer/db"
	"pentag.kr/distimer/ent/affiliation"
	"pentag.kr/distimer/middlewares"
	"pentag.kr/distimer/utils/logger"
)

// @Summary Get All Group Members
// @Tags Group
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "group id"
// @Success 200 {array} AffiliationDTO
// @Router /group/member/{id} [get]
func GetAllGroupMembers(c *fiber.Ctx) error {
	groupIDStr := c.Params("id")
	groupID, err := uuid.Parse(groupIDStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid group id",
		})
	}
	userID := middlewares.GetUserIDFromMiddleware(c)

	dbConn := db.GetDBClient()

	userExistInGroup, err := dbConn.Affiliation.Query().Where(affiliation.And(affiliation.GroupID(groupID), affiliation.UserID(userID))).Exist(context.Background())
	if err != nil {
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	} else if !userExistInGroup {
		return c.Status(404).JSON(fiber.Map{
			"error": "Group not found or you are not the member of the group",
		})
	}

	affiliations, err := dbConn.Affiliation.Query().Where(affiliation.GroupID(groupID)).All(context.Background())
	if err != nil {
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	result := make([]AffiliationDTO, len(affiliations))
	for i, affiliation := range affiliations {
		result[i] = AffiliationDTO{
			UserID:   affiliation.UserID.String(),
			GroupID:  affiliation.GroupID.String(),
			Nickname: affiliation.Nickname,
			Role:     affiliation.Role,
			JoinedAt: affiliation.JoinedAt.Format(time.RFC3339),
		}
	}
	return c.JSON(result)
}
