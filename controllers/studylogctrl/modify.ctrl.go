package studylogctrl

import (
	"context"
	"time"
	"unicode/utf8"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"pentag.kr/distimer/configs"
	"pentag.kr/distimer/db"
	"pentag.kr/distimer/ent"
	"pentag.kr/distimer/ent/affiliation"
	"pentag.kr/distimer/ent/studylog"
	"pentag.kr/distimer/ent/subject"
	"pentag.kr/distimer/ent/user"
	"pentag.kr/distimer/middlewares"
	"pentag.kr/distimer/utils/dto"
	"pentag.kr/distimer/utils/logger"
)

// @Summary Modify Study Log
// @Tags StudyLog
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "studylog id"
// @Param request body createStudyLogReq true "createStudyLogReq"
// @Success 200 {object} createStudyLogReq
// @Failure 400
// @Failure 403
// @Failure 404
// @Failure 409
// @Failure 500
// @Router /studylog/{id} [put]
func ModifyStudyLog(c *fiber.Ctx) error {

	studylogIDStr := c.Params("id")
	studylogID, err := uuid.Parse(studylogIDStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid studylog id",
		})
	}

	// Data Parsing
	userID := middlewares.GetUserIDFromMiddleware(c)

	data := new(createStudyLogReq)
	if err := dto.Bind(c, data); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err,
		})
	}
	if utf8.RuneCountInString(data.Content) > 30 {
		return c.Status(400).JSON(fiber.Map{
			"error": "Content length should be between 1 and 30",
		})
	}
	// parse date with rfc3339 format
	startAt, err := time.Parse(time.RFC3339, data.StartAt)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid start_at",
		})
	}
	endAt, err := time.Parse(time.RFC3339, data.EndAt)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid end_at",
		})
	}
	if startAt.After(endAt) {
		return c.Status(400).JSON(fiber.Map{
			"error": "start_at should be before end_at",
		})
	}

	subjectID, err := uuid.Parse(data.SubjectID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid subject ID",
		})
	}

	startAt = startAt.Truncate(time.Second)
	endAt = endAt.Truncate(time.Second)

	var groupIDs []uuid.UUID
	count := 1
	for _, groupIDStr := range data.GroupsToShare {
		if count > configs.FreePlanGroupJoinLimit {
			return c.Status(400).JSON(fiber.Map{
				"error": "You can't share more than limit of free plan groups",
			})
		}

		id, err := uuid.Parse(groupIDStr)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Invalid group ID",
			})
		}
		groupIDs = append(groupIDs, id)

		count++
	}

	dbConn := db.GetDBClient()

	studylogObj, err := dbConn.StudyLog.Query().
		Where(
			studylog.And(
				studylog.ID(studylogID),
				studylog.HasUserWith(user.ID(userID))),
		).WithSubject().WithSharedGroup().Only(context.Background())
	if err != nil {
		if ent.IsNotFound(err) {
			return c.Status(404).JSON(fiber.Map{
				"error": "Studylog not found or you are not the owner",
			})
		}
		logger.CtxError(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	// check if study log is already exist at the same time
	duplicatedStudyLogs, err := dbConn.StudyLog.Query().Where(studylog.And(studylog.StartAtLTE(endAt), studylog.EndAtGTE(startAt))).All(context.Background())
	if err != nil {
		logger.CtxError(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	for _, studylogObj := range duplicatedStudyLogs {
		if studylogObj.ID != studylogID {
			return c.Status(409).JSON(fiber.Map{
				"error": "Study log is already exist at the same time",
			})
		}
	}

	// check if the subject is owned by the user
	if studylogObj.Edges.Subject.ID != subjectID {
		subjectObj, err := dbConn.Subject.Query().Where(subject.ID(subjectID)).WithCategory().Only(context.Background())
		if err != nil {
			if ent.IsNotFound(err) {
				return c.Status(404).JSON(fiber.Map{
					"error": "Subject not found",
				})
			}
			logger.CtxError(c, err)
			return c.Status(500).JSON(fiber.Map{
				"error": "Internal server error",
			})
		}

		userObj, err := subjectObj.Edges.Category.QueryUser().Only(context.Background())
		if err != nil {
			logger.CtxError(c, err)
			return c.Status(500).JSON(fiber.Map{
				"error": "Internal server error",
			})
		}
		if userObj.ID != userID {
			return c.Status(403).JSON(fiber.Map{
				"error": "You are not the owner of the subject",
			})
		}
	}

	// check if the user is the member of the group
	for _, groupID := range groupIDs {
		exist, err := dbConn.Affiliation.Query().Where(affiliation.And(affiliation.GroupID(groupID), affiliation.UserID(userID))).Exist(context.Background())
		if err != nil {
			logger.CtxError(c, err)
			return c.Status(500).JSON(fiber.Map{
				"error": "Internal server error",
			})
		} else if !exist {
			return c.Status(403).JSON(fiber.Map{
				"error": "You are not the member of the group: " + groupID.String(),
			})
		}
	}

	_, err = dbConn.StudyLog.UpdateOne(studylogObj).
		SetContent(data.Content).
		SetStartAt(startAt).
		SetEndAt(endAt).
		SetSubjectID(subjectID).
		ClearSharedGroup().
		AddSharedGroupIDs(groupIDs...).
		Save(context.Background())
	if err != nil {
		logger.CtxError(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	return c.JSON(data)
}
