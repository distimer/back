package categoryctrl

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"pentag.kr/distimer/db"
	"pentag.kr/distimer/ent"
	"pentag.kr/distimer/ent/category"
	"pentag.kr/distimer/ent/user"
	"pentag.kr/distimer/middlewares"
	"pentag.kr/distimer/utils/logger"
)

// @Summary Delete Category
// @Tags Category
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "Category ID"
// @Success 204
// @Failure 400
// @Failure 404
// @Failure 409
// @Failure 500
// @Router /category/{id} [delete]
func DeleteCategory(c *fiber.Ctx) error {
	categoryIDStr := c.Params("id")
	categoryID, err := uuid.Parse(categoryIDStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid category ID",
		})
	}

	userID := middlewares.GetUserIDFromMiddleware(c)

	dbConn := db.GetDBClient()
	categoryObj, err := dbConn.Category.Query().Where(category.And(category.ID(categoryID), category.HasUserWith(user.ID(userID)))).WithSubjects().Only(context.Background())
	if err != nil {
		if ent.IsNotFound(err) {
			return c.Status(404).JSON(fiber.Map{
				"error": "Category not found",
			})
		}
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}
	if len(categoryObj.Edges.Subjects) != 0 {
		return c.Status(409).JSON(fiber.Map{
			"error": "Category has subjects",
		})
	}

	err = dbConn.Category.DeleteOne(categoryObj).Exec(context.Background())
	if err != nil {
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	return c.SendStatus(204)
}
