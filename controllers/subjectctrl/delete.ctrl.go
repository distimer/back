package subjectctrl

import (
	"context"
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"pentag.kr/distimer/db"
	"pentag.kr/distimer/ent"
	"pentag.kr/distimer/ent/category"
	"pentag.kr/distimer/ent/studylog"
	"pentag.kr/distimer/ent/subject"
	"pentag.kr/distimer/ent/user"
	"pentag.kr/distimer/middlewares"
	"pentag.kr/distimer/utils/logger"
)

// @Summary Delete Subject
// @Tags Subject
// @Security Bearer
// @Param id path string true "Subject ID"
// @Success 204
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /subject/{id} [delete]
func DeleteSubject(c *fiber.Ctx) error {
	subjectIDStr := c.Params("id")
	subjectID, err := uuid.Parse(subjectIDStr)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid subject ID",
		})
	}

	userID := middlewares.GetUserIDFromMiddleware(c)

	dbConn := db.GetDBClient()

	subjectObj, err := dbConn.Subject.Get(context.Background(), subjectID)
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

	// Get the category object
	categoryObj, err := subjectObj.QueryCategory().Only(context.Background())
	if err != nil {
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}
	userObj, err := categoryObj.QueryUser().Only(context.Background())
	if err != nil {
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}
	if userObj.ID != userID {
		return c.Status(404).JSON(fiber.Map{
			"error": "Subject not found",
		})
	} else if categoryObj.Name == "미분류" {
		return c.Status(403).JSON(fiber.Map{
			"error": "You cannot delete the default subject",
		})
	}

	defaultCategory, err := dbConn.Category.Query().
		Where(category.And(category.Name("미분류"), category.HasUserWith(user.ID(userID)))).
		WithSubjects().
		Only(context.Background())
	if err != nil {
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	} else if len(defaultCategory.Edges.Subjects) != 1 {
		logger.Error(c, errors.New("default category should have exactly one subject"))
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	err = dbConn.StudyLog.Update().Where(studylog.HasSubjectWith(subject.ID(subjectID))).SetSubject(defaultCategory.Edges.Subjects[0]).Exec(context.Background())
	if err != nil {
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	err = dbConn.Subject.DeleteOne(subjectObj).Exec(context.Background())
	if err != nil {
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}
	return c.SendStatus(204)
}
