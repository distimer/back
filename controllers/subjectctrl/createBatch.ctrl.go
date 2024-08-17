package subjectctrl

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

type createSubjectBatchElemReq struct {
	Name       string `json:"name" validate:"required" example:"name between 1 and 20"`
	Color      string `json:"color" validate:"required,hexcolor"`
	CategoryID string `json:"category_id" validate:"required,uuid"`
}

type createSubjectBatchReq struct {
	SubjectList []createSubjectBatchElemReq `json:"subject_list" validate:"required"`
}

// @Summary Create Batch Subject
// @Tags Subject
// @Accept json
// @Produce json
// @Security Bearer
// @Param data body createSubjectBatchReq true "subject data"
// @Success 200 {array} SubjectDTO
// @Failure 400
// @Failure 404
// @Failure 409
// @Failure 500
// @Router /subject/batch [post]
func CreateBatchSubject(c *fiber.Ctx) error {
	data := new(createSubjectBatchReq)
	if err := dto.Bind(c, data); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err,
		})
	}
	for _, subject := range data.SubjectList {
		if utf8.RuneCountInString(subject.Name) < 1 || utf8.RuneCountInString(subject.Name) > 20 {
			return c.Status(400).JSON(fiber.Map{
				"error": "Name length should be between 1 and 20",
			})
		} else if subject.Name == "미분류" {
			return c.Status(400).JSON(fiber.Map{
				"error": "Cannot add subject with default name",
			})
		}
	}
	userID := middlewares.GetUserIDFromMiddleware(c)
	dbConn := db.GetDBClient()
	for _, subject := range data.SubjectList {
		categoryID := uuid.MustParse(subject.CategoryID)
		_, err := dbConn.Category.Query().
			Where(category.And(category.ID(categoryID), category.HasUserWith(user.ID(userID)))).
			Only(context.Background())
		if err != nil {
			if ent.IsNotFound(err) {
				return c.Status(400).JSON(fiber.Map{
					"error": "Category not found",
				})
			}
			logger.CtxError(c, err)
			return c.Status(500).JSON(fiber.Map{
				"error": "Internal server error",
			})
		}
	}
	var result []SubjectDTO
	for _, subject := range data.SubjectList {
		categoryID := uuid.MustParse(subject.CategoryID)
		newSubject, err := dbConn.Subject.Create().
			SetName(subject.Name).
			SetColor(subject.Color).
			SetCategoryID(categoryID).
			Save(context.Background())
		if err != nil {
			logger.CtxError(c, err)
			return c.Status(500).JSON(fiber.Map{
				"error": "Internal server error",
			})
		}
		result = append(result, SubjectDTO{
			ID:    newSubject.ID.String(),
			Name:  newSubject.Name,
			Color: newSubject.Color,
			Order: newSubject.Order,
		})
	}
	return c.JSON(result)
}
