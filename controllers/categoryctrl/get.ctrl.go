package categoryctrl

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"pentag.kr/distimer/controllers/subjectctrl"
	"pentag.kr/distimer/db"
	"pentag.kr/distimer/ent"
	"pentag.kr/distimer/ent/category"
	"pentag.kr/distimer/ent/user"
	"pentag.kr/distimer/middlewares"
)

// @Summary Get Category List
// @Tags Category
// @Description [EDGE INCLUDED!]Subject list is included in each category
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {array} categoryDTO
// @Router /category [get]
func GetCategoryList(c *fiber.Ctx) error {
	userID := middlewares.GetUserIDFromMiddleware(c)

	dbConn := db.GetDBClient()

	var err error
	categories, err := dbConn.Category.Query().Where(category.HasUserWith(user.ID(userID))).WithSubjects().All(context.Background())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}
	if categories == nil {
		categories = []*ent.Category{}
	}

	result := make([]categoryDTO, len(categories))
	for i, category := range categories {
		result[i] = categoryDTO{
			ID:    category.ID.String(),
			Name:  category.Name,
			Order: category.Order,
			Subjects: func() []subjectctrl.SubjectDTO {
				subjects := category.Edges.Subjects
				if subjects == nil {
					subjects = []*ent.Subject{}
				}

				result := make([]subjectctrl.SubjectDTO, len(subjects))
				for i, subject := range subjects {
					result[i] = subjectctrl.SubjectDTO{
						ID:    subject.ID.String(),
						Name:  subject.Name,
						Color: subject.Color,
						Order: subject.Order,
					}
				}

				return result
			}(),
		}
	}

	return c.JSON(result)
}
