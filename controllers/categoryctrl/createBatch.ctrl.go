package categoryctrl

import (
	"context"
	"unicode/utf8"

	"github.com/gofiber/fiber/v2"
	"pentag.kr/distimer/configs"
	"pentag.kr/distimer/controllers/subjectctrl"
	"pentag.kr/distimer/db"
	"pentag.kr/distimer/ent/category"
	"pentag.kr/distimer/ent/user"
	"pentag.kr/distimer/middlewares"
	"pentag.kr/distimer/utils/dto"
	"pentag.kr/distimer/utils/logger"
)

type createCategoryBatchReq struct {
	CategoryList []string `json:"category_list" validate:"required"`
}

// @Summary Create Batch Category
// @Tags Category
// @Accept json
// @Produce json
// @Security Bearer
// @Param data body createCategoryBatchReq true "category data"
// @Success 200 {array} categoryDTO
// @Failure 400
// @Failure 409
// @Failure 500
// @Router /category/batch [post]
func CreateBatchCategory(c *fiber.Ctx) error {
	data := new(createCategoryBatchReq)
	if err := dto.Bind(c, data); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err,
		})
	}
	for _, categoryName := range data.CategoryList {
		if utf8.RuneCountInString(categoryName) < 1 || utf8.RuneCountInString(categoryName) > 20 {
			return c.Status(400).JSON(fiber.Map{
				"error": "Name length should be between 1 and 20",
			})
		} else if categoryName == "미분류" {
			return c.Status(400).JSON(fiber.Map{
				"error": "Cannot add category with default name",
			})
		}
	}

	userID := middlewares.GetUserIDFromMiddleware(c)

	dbConn := db.GetDBClient()

	userCategoryCnt, err := dbConn.Category.Query().Where(category.HasUserWith(user.ID(userID))).Count(context.Background())
	if err != nil {
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}
	if userCategoryCnt+len(data.CategoryList) > configs.FreePlanCategoryLimit {
		return c.Status(409).JSON(fiber.Map{
			"error": "Category limit exceeded",
		})
	}

	result := []categoryDTO{}
	for _, categoryName := range data.CategoryList {
		categoryObj, err := dbConn.Category.Create().
			SetName(categoryName).
			SetUserID(userID).
			Save(context.Background())
		if err != nil {
			logger.Error(c, err)
			return c.Status(500).JSON(fiber.Map{
				"error": "Internal server error",
			})
		}
		result = append(result, categoryDTO{
			ID:       categoryObj.ID.String(),
			Name:     categoryObj.Name,
			Order:    categoryObj.Order,
			Subjects: []subjectctrl.SubjectDTO{},
		})

	}

	return c.JSON(result)
}
