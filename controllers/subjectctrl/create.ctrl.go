package subjectctrl

import (
	"context"
	"unicode/utf8"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"pentag.kr/distimer/db"
	"pentag.kr/distimer/ent/category"
	"pentag.kr/distimer/ent/user"
	"pentag.kr/distimer/middlewares"
	"pentag.kr/distimer/utils/dto"
	"pentag.kr/distimer/utils/logger"
)

type createSubjectReq struct {
	Name  string `json:"name" validate:"required" example:"name between 1 and 20"`
	Color string `json:"color" validate:"required,hexcolor"`
}

type createSubjectRes struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

// @Summary Create Subject
// @Tags Subject
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "Category ID"
// @Param request body createSubjectReq true "createSubjectReq"
// @Success 200 {object} createSubjectRes
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /subject/{id} [post]
func CreateSubject(c *fiber.Ctx) error {
	categoryIDStr := c.Params("id")
	categoryID, err := uuid.Parse(categoryIDStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid category ID",
		})
	}

	data := new(createSubjectReq)
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

	categoryExist, err := dbConn.Category.Query().Where(category.And(category.ID(categoryID), category.HasUserWith(user.ID(userID)))).Exist(context.Background())
	if err != nil {
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	} else if !categoryExist {
		return c.Status(404).JSON(fiber.Map{
			"error": "Category not found or you are not the owner",
		})
	}

	subjectObj, err := dbConn.Subject.Create().
		SetName(data.Name).
		SetColor(data.Color).
		SetCategoryID(categoryID).
		Save(context.Background())
	if err != nil {
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	return c.JSON(
		createSubjectRes{
			ID:    subjectObj.ID.String(),
			Name:  subjectObj.Name,
			Color: subjectObj.Color,
		},
	)
}
