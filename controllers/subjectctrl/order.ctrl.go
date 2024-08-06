package subjectctrl

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"pentag.kr/distimer/configs"
	"pentag.kr/distimer/db"
	"pentag.kr/distimer/ent/subject"
	"pentag.kr/distimer/middlewares"
	"pentag.kr/distimer/utils/dto"
	"pentag.kr/distimer/utils/logger"
)

type subjectOrderElement struct {
	SubjectID string `json:"subject_id" validate:"required,uuid"`
	Order     int8   `json:"order" validate:"required,min=0"`
}

// @Summary Modify Subject Order
// @Tags Subject
// @Accept json
// @Produce json
// @Security Bearer
// @Param request body []subjectOrderElement true "subjectOrderElement"
// @Success 200 {object} []subjectOrderElement
// @Failure 400
// @Failure 403
// @Failure 500
// @Router /subject/order [patch]
func SubjectOrderModify(c *fiber.Ctx) error {
	data := new([]subjectOrderElement)
	if err := dto.Bind(c, data); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err,
		})
	}
	if len(*data) == 0 || len(*data) > configs.FreePlanCategoryLimit*configs.FreePlanSubjectPerCategoryLimit {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid number of elements",
		})
	}
	userID := middlewares.GetUserIDFromMiddleware(c)
	dbConn := db.GetDBClient()
	for _, element := range *data {
		subjectID := uuid.MustParse(element.SubjectID)

		// Check if the subject exists and the user has permission to modify the order

		subjectObj, err := dbConn.Subject.Query().Where(subject.ID(subjectID)).WithCategory().Only(context.Background())
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Invalid subject ID" + element.SubjectID,
			})
		}
		userObj, err := subjectObj.Edges.Category.QueryUser().Only(context.Background())
		if err != nil {
			logger.Error(c, err)
			return c.Status(500).JSON(fiber.Map{
				"error": "Internal server error",
			})
		}

		if userObj.ID != userID {
			return c.Status(403).JSON(fiber.Map{
				"error": "You are not the owner of the subject",
			})
		}

		// Update the order
		if _, err = dbConn.Subject.UpdateOne(subjectObj).SetOrder(element.Order).Save(context.Background()); err != nil {
			logger.Error(c, err)
			return c.Status(500).JSON(fiber.Map{
				"error": "Internal server error",
			})
		}
	}
	return c.JSON(data)
}
