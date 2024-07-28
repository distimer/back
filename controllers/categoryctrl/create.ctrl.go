package categoryctrl

import (
	"context"
	"unicode/utf8"

	"github.com/gofiber/fiber/v2"
	"pentag.kr/distimer/db"
	"pentag.kr/distimer/ent"
	"pentag.kr/distimer/middlewares"
	"pentag.kr/distimer/utils/dto"
	"pentag.kr/distimer/utils/logger"
)

type createCategoryReq struct {
	Name string `json:"name" validate:"required" example:"name between 1 and 20"`
}

type createCategoryRes struct {
	Category *ent.Category `json:"category"`
}

// @Summary Create Category
// @Tags Category
// @Accept json
// @Produce json
// @Security Bearer
// @Param request body createCategoryReq true "createCategoryReq"
// @Success 200 {object} createCategoryRes
// @Router /category [post]
func CreateCategory(c *fiber.Ctx) error {
	data := new(createCategoryReq)
	if err := dto.Bind(c, data); err != nil {
		return err
	}
	if utf8.RuneCountInString(data.Name) < 1 || utf8.RuneCountInString(data.Name) > 20 {
		return c.Status(400).JSON(fiber.Map{
			"error": "Name length should be between 1 and 20",
		})
	}

	userID := middlewares.GetUserIDFromMiddleware(c)

	dbConn := db.GetDBClient()

	categoryObj, err := dbConn.Category.Create().
		SetName(data.Name).
		SetUserID(userID).
		Save(context.Background())
	if err != nil {
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	return c.JSON(
		createCategoryRes{
			Category: categoryObj,
		},
	)
}
