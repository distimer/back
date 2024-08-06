package categoryctrl

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"pentag.kr/distimer/configs"
	"pentag.kr/distimer/db"
	"pentag.kr/distimer/ent"
	"pentag.kr/distimer/ent/category"
	"pentag.kr/distimer/middlewares"
	"pentag.kr/distimer/utils/dto"
	"pentag.kr/distimer/utils/logger"
)

type categoryOrderElement struct {
	CategoryID string `json:"category_id" validate:"required,uuid"`
	Order      int8   `json:"order" validate:"required,min=0"`
}

// @Summary Modify Category Order
// @Tags Category
// @Accept json
// @Produce json
// @Security Bearer
// @Param request body []categoryOrderElement true "categoryOrderElement"
// @Success 200 {object} []categoryOrderElement
// @Failure 400
// @Failure 403
// @Failure 500
// @Router /category/order [patch]
func CategoryOrderModify(c *fiber.Ctx) error {
	data := new([]categoryOrderElement)

	if err := dto.Bind(c, data); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err,
		})
	}

	if len(*data) == 0 || len(*data) > configs.FreePlanCategoryLimit {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid number of elements",
		})
	}

	userID := middlewares.GetUserIDFromMiddleware(c)
	dbConn := db.GetDBClient()
	for _, element := range *data {
		// Check if the category exists and the user has permission to modify the order
		categoryID := uuid.MustParse(element.CategoryID)
		categoryObj, err := dbConn.Category.Query().Where(category.ID(categoryID)).WithUser().Only(context.Background())
		if err != nil {
			if ent.IsNotFound(err) {
				return c.Status(400).JSON(fiber.Map{
					"error": "Invalid category ID",
				})
			}
			logger.Error(c, err)
			return c.Status(500).JSON(fiber.Map{
				"error": "Internal Server Error",
			})
		}
		if categoryObj.Edges.User.ID != userID {
			return c.Status(403).JSON(fiber.Map{
				"error": "You are not the owner of the category",
			})
		}

		// Update the order
		if _, err = dbConn.Category.UpdateOne(categoryObj).SetOrder(element.Order).Save(context.Background()); err != nil {
			logger.Error(c, err)
			return c.Status(500).JSON(fiber.Map{
				"error": "Internal Server Error",
			})
		}

	}
	return c.JSON(data)
}
