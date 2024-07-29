package subjectctrl

import (
	"context"
	"unicode/utf8"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"pentag.kr/distimer/db"
	"pentag.kr/distimer/ent"
	"pentag.kr/distimer/ent/subject"
	"pentag.kr/distimer/middlewares"
	"pentag.kr/distimer/utils/dto"
	"pentag.kr/distimer/utils/logger"
)

type modifySubjectInfoReq struct {
	Name  string `json:"name" validate:"required" example:"name between 1 and 20"`
	Color string `json:"color" validate:"required,hexcolor"`
}

// @Summary Modify Subject Info
// @Tags Subject
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "Subject ID"
// @Param request body modifySubjectInfoReq true "modifySubjectInfoReq"
// @Success 200 {object} SubjectDTO
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /subject/{id} [put]
func ModifySubjectInfo(c *fiber.Ctx) error {
	subjectIDStr := c.Params("id")
	subjectID, err := uuid.Parse(subjectIDStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid subject ID",
		})
	}

	data := new(modifySubjectInfoReq)
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
	subjectObj, err := dbConn.Subject.Query().Where(subject.ID(subjectID)).WithCategory().Only(context.Background())

	if err != nil {
		if ent.IsNotFound(err) {
			return c.Status(404).JSON(fiber.Map{
				"error": "Subject not found",
			})
		}
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})

	}
	if subjectObj.Edges.Category.Edges.User.ID != userID {
		return c.Status(404).JSON(fiber.Map{
			"error": "Subject not found",
		})
	}

	_, err = subjectObj.Update().SetName(data.Name).SetColor(data.Color).Save(context.Background())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	return c.JSON(SubjectDTO{
		ID:    subjectID.String(),
		Name:  data.Name,
		Color: data.Color,
	})
}
