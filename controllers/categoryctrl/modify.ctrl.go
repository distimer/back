package categoryctrl

import (
	"context"
	"unicode/utf8"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"pentag.kr/distimer/db"
	"pentag.kr/distimer/ent"
	"pentag.kr/distimer/ent/category"
	"pentag.kr/distimer/ent/user"
	"pentag.kr/distimer/middlewares"
	"pentag.kr/distimer/utils/dto"
	"pentag.kr/distimer/utils/logger"
)

type modifyCategoryRes struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// @Summary Modify Category
// @Tags Category
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "Category ID"
// @Param request body createCategoryReq true "createCategoryReq"
// @Success 200 {object} modifyCategoryRes
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /category/{id} [put]
func ModifyCategory(c *fiber.Ctx) error {

	categoryIDStr := c.Params("id")
	categoryID, err := uuid.Parse(categoryIDStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid category ID",
		})
	}

	data := new(createCategoryReq)
	if err := dto.Bind(c, data); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err,
		})
	}
	if utf8.RuneCountInString(data.Name) < 1 || utf8.RuneCountInString(data.Name) > 20 {
		return c.Status(400).JSON(fiber.Map{
			"error": "Name length should be between 1 and 20",
		})
	}

	userID := middlewares.GetUserIDFromMiddleware(c)

	dbConn := db.GetDBClient()

	categoryObj, err := dbConn.Category.Query().
		Where(category.And(category.ID(categoryID), category.HasUserWith(user.ID(userID)))).
		Only(context.Background())

	if err != nil {
		if ent.IsNotFound(err) {
			return c.Status(404).JSON(fiber.Map{
				"error": "Category not found or you are not the owner",
			})
		}
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	_, err = categoryObj.Update().
		SetName(data.Name).
		Save(context.Background())
	if err != nil {
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	return c.JSON(modifyCategoryRes{
		ID:   categoryObj.ID.String(),
		Name: data.Name,
	})
}
