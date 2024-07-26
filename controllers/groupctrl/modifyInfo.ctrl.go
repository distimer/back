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

type modifyGroupInfoReq struct {
	Name        string `json:"name" validate:"required" example:"name between 3 and 30"`
	Description string `json:"description" example:"description between 0 and 100"`
}

type modifyGroupInfoRes struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// @Summary Modify Group Info
// @Tags Group
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "group id"
// @Param request body groupctrl.modifyGroupInfoReq true "modifyGroupInfoReq"
// @Success 200 {object} modifyGroupInfoRes
// @Router /group/{id} [put]
func ModifyGroupInfo(c *fiber.Ctx) error {
	groupIDStr := c.Params("id")
	groupID, err := uuid.Parse(groupIDStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid group id",
		})
	}

	data := new(modifyGroupInfoReq)
	if err := dto.Bind(c, data); err != nil {
		return err
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

	err = groupObj.Update().SetName(data.Name).SetDescription(data.Description).Exec(context.Background())
	if err != nil {
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	return c.JSON(modifyGroupInfoRes{
		ID:          groupID.String(),
		Name:        data.Name,
		Description: data.Description,
	})
}
