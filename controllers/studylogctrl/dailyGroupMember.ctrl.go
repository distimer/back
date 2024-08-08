package studylogctrl

// import (
// 	"context"
// 	"time"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/google/uuid"
// 	"pentag.kr/distimer/controllers/subjectctrl"
// 	"pentag.kr/distimer/db"
// 	"pentag.kr/distimer/ent"
// 	"pentag.kr/distimer/ent/affiliation"
// 	"pentag.kr/distimer/ent/group"
// 	"pentag.kr/distimer/ent/studylog"
// 	"pentag.kr/distimer/middlewares"
// 	"pentag.kr/distimer/utils/logger"
// )

// type dailyGroupMemberLogElem struct {
// 	Subject   subjectctrl.SubjectDTO `json:"subject" validate:"required"`
// 	StudyTime int                    `json:"study_time" validate:"required"`
// }

// type dailyGroupMemberLogResponse struct {
// 	UserID uuid.UUID                 `json:"user_id" validate:"required"`
// 	Log    []dailyGroupMemberLogElem `json:"log" validate:"required"`
// }

// func DailyGroupMemberLog(c *fiber.Ctx) error {

// 	dateStr := c.Query("date")
// 	date, err := time.Parse("2006-01-02", dateStr)
// 	if err != nil {
// 		return c.Status(400).JSON(fiber.Map{
// 			"error": "Invalid date format",
// 		})
// 	}

// 	groupIDStr := c.Params("id")
// 	groupID, err := uuid.Parse(groupIDStr)
// 	if err != nil {
// 		return c.Status(400).JSON(fiber.Map{
// 			"error": "Invalid group id",
// 		})
// 	}

// 	userID := middlewares.GetUserIDFromMiddleware(c)

// 	dbConn := db.GetDBClient()

// 	groupObj, err := dbConn.Group.Query().Where(group.ID(groupID)).WithOwner().Only(context.Background())
// 	if err != nil {
// 		if ent.IsNotFound(err) {
// 			return c.Status(404).JSON(fiber.Map{
// 				"error": "Group not found",
// 			})
// 		}
// 		logger.Error(c, err)
// 		return c.Status(500).JSON(fiber.Map{
// 			"error": "Internal server error",
// 		})
// 	}

// 	if groupObj.Edges.Owner.ID != userID {
// 		userAffilitaionObj, err := dbConn.Affiliation.Query().Where(
// 			affiliation.And(
// 				affiliation.GroupID(groupID),
// 				affiliation.UserID(userID),
// 			),
// 		).Only(context.Background())
// 		if err != nil {
// 			if ent.IsNotFound(err) {
// 				return c.Status(404).JSON(fiber.Map{
// 					"error": "You are not the member of the group",
// 				})
// 			}
// 			logger.Error(c, err)
// 			return c.Status(500).JSON(fiber.Map{
// 				"error": "Internal server error",
// 			})
// 		}
// 		if userAffilitaionObj.Role < groupObj.RevealPolicy {
// 			return c.Status(403).JSON(fiber.Map{
// 				"error": "You are not allowed to see the log",
// 			})
// 		}
// 	}
// 	// get all members
// 	members, err := dbConn.Affiliation.Query().Where(affiliation.GroupID(groupID)).All(context.Background())
// 	if err != nil {
// 		logger.Error(c, err)
// 		return c.Status(500).JSON(fiber.Map{
// 			"error": "Internal server error",
// 		})
// 	}

// 	// get all shared study logs
// 	studyLogList, err := dbConn.StudyLog.Query().Where(
// 		studylog.And(
// 			studylog.StartAtLTE(date.AddDate(0, 0, 1)),
// 			studylog.EndAtGTE(date),
// 			studylog.HasSharedGroupWith(group.ID(groupID)),
// 		),
// 	).WithUser().WithSubject().All(context.Background())
// 	if err != nil {
// 		logger.Error(c, err)
// 		return c.Status(500).JSON(fiber.Map{
// 			"error": "Internal server error",
// 		})
// 	}

// 	result := make([]dailyGroupMemberLogResponse, len(members))
// 	for i, member := range members {

// 	}

// 	return nil
// }
