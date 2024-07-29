package userctrl

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"pentag.kr/distimer/db"
	"pentag.kr/distimer/ent"
	"pentag.kr/distimer/ent/affiliation"
	"pentag.kr/distimer/ent/refreshtoken"
	"pentag.kr/distimer/ent/studylog"
	"pentag.kr/distimer/ent/user"
	"pentag.kr/distimer/middlewares"
	"pentag.kr/distimer/utils/logger"
)

// @Summary Quit Service
// @Tags User
// @Security Bearer
// @Success 204
// @Failure 409
// @Failure 500
// @Router /user [delete]
func QuitService(c *fiber.Ctx) error {
	userID := middlewares.GetUserIDFromMiddleware(c)

	dbConn := db.GetDBClient()

	foundUser, err := dbConn.User.Get(context.Background(), userID)
	if err != nil {
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	// group owner check
	ownedGroups, err := foundUser.QueryOwnedGroups().All(context.Background())
	if err != nil {
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}
	if len(ownedGroups) != 0 {
		return c.Status(409).JSON(fiber.Map{
			"error": "Group owner cannot quit service",
		})
	}

	// refresh token deletion
	_, err = dbConn.RefreshToken.Delete().Where(refreshtoken.HasUserWith(user.ID(userID))).Exec(context.Background())
	if err != nil {
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	// timer deletion
	userTimer, err := foundUser.QueryTimers().Only(context.Background())
	if err != nil && !ent.IsNotFound(err) {
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	} else if err == nil && !ent.IsNotFound(err) {
		err = dbConn.Timer.DeleteOne(userTimer).Exec(context.Background())
		if err != nil {
			logger.Error(c, err)
			return c.Status(500).JSON(fiber.Map{
				"error": "Internal server error",
			})
		}
	}

	// study log deletion
	_, err = dbConn.StudyLog.Delete().Where(studylog.HasUserWith(user.ID(userID))).Exec(context.Background())
	if err != nil {
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	// category and subject deletion
	categories, err := foundUser.QueryOwnedCategories().All(context.Background())
	if err != nil {
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}
	for _, category := range categories {
		subjects, err := category.QuerySubjects().All(context.Background())
		if err != nil {
			logger.Error(c, err)
			return c.Status(500).JSON(fiber.Map{
				"error": "Internal server error",
			})
		}
		for _, subject := range subjects {
			err = dbConn.Subject.DeleteOne(subject).Exec(context.Background())
			if err != nil {
				logger.Error(c, err)
				return c.Status(500).JSON(fiber.Map{
					"error": "Internal server error",
				})
			}
		}
		err = dbConn.Category.DeleteOne(category).Exec(context.Background())
		if err != nil {
			logger.Error(c, err)
			return c.Status(500).JSON(fiber.Map{
				"error": "Internal server error",
			})
		}
	}

	// affiliation deletion
	_, err = dbConn.Affiliation.Delete().Where(affiliation.UserID(userID)).Exec(context.Background())
	if err != nil {
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	// user deletion
	err = dbConn.User.DeleteOne(foundUser).Exec(context.Background())
	if err != nil {
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	return c.Status(204).JSON(fiber.Map{
		"message": "Bye Bye :(",
	})
}